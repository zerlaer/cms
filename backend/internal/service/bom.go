package service

import (
	"cms/internal/models"
	"errors"
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type BOMService struct {
	db *gorm.DB
}

func NewBOMService(db *gorm.DB) *BOMService {
	return &BOMService{db: db}
}

// BOMItem BOM项目
type BOMItem struct {
	Comment    string            `json:"comment"`
	Footprint  string            `json:"footprint"`
	Quantity   int               `json:"quantity"`
	Model      string            `json:"model"`
	Parameters string            `json:"parameters"`
	Matched    bool              `json:"matched"`
	Component  *models.Component `json:"component,omitempty"`
	MatchScore int               `json:"matchScore"`
}

// ImportBOM 导入BOM文件
func (s *BOMService) ImportBOM(filePath string) ([]BOMItem, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 获取第一个工作表
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) < 2 {
		return nil, errors.New("BOM file has no data")
	}

	// 解析表头
	header := make(map[string]int)
	for i, cell := range rows[0] {
		header[strings.TrimSpace(cell)] = i
	}

	// 验证必需的列
	requiredColumns := []string{"Comment", "Footprint", "Quantity"}
	for _, col := range requiredColumns {
		if _, exists := header[col]; !exists {
			return nil, fmt.Errorf("missing required column: %s", col)
		}
	}

	// 解析数据
	var bomItems []BOMItem
	for i, row := range rows {
		if i == 0 { // 跳过表头
			continue
		}

		item := BOMItem{
			Comment:    getCellValue(row, header, "Comment"),
			Footprint:  getCellValue(row, header, "Footprint"),
			Model:      getCellValue(row, header, "型号"),
			Parameters: getCellValue(row, header, "参数"),
		}

		// 解析数量
		quantityStr := getCellValue(row, header, "Quantity")
		if quantityStr != "" {
			var quantity int
			if _, err := fmt.Sscanf(quantityStr, "%d", &quantity); err == nil {
				item.Quantity = quantity
			}
		}

		bomItems = append(bomItems, item)
	}

	return bomItems, nil
}

// MatchBOM 匹配BOM到库存
func (s *BOMService) MatchBOM(bomItems []BOMItem) ([]BOMItem, error) {
	// 获取所有元件
	var components []models.Component
	if err := s.db.Find(&components).Error; err != nil {
		return nil, err
	}

	for i := range bomItems {
		bestMatch := s.findBestMatch(&bomItems[i], components)
		if bestMatch != nil {
			bomItems[i].Matched = true
			bomItems[i].Component = bestMatch
			// 计算匹配分数
			bomItems[i].MatchScore = calculateMatchScore(&bomItems[i], bestMatch)
		}
	}

	return bomItems, nil
}

// findBestMatch 查找最佳匹配的元件
func (s *BOMService) findBestMatch(bomItem *BOMItem, components []models.Component) *models.Component {
	var bestMatch *models.Component
	bestScore := 0

	for i := range components {
		score := calculateMatchScore(bomItem, &components[i])
		if score > bestScore {
			bestScore = score
			bestMatch = &components[i]
		}
	}

	// 只有分数大于阈值才算匹配
	if bestScore > 50 {
		return bestMatch
	}
	return nil
}

// calculateMatchScore 计算匹配分数
func calculateMatchScore(bomItem *BOMItem, component *models.Component) int {
	score := 0

	// 匹配Comment
	if strings.Contains(strings.ToLower(component.Name), strings.ToLower(bomItem.Comment)) ||
		strings.Contains(strings.ToLower(bomItem.Comment), strings.ToLower(component.Name)) {
		score += 30
	}

	// 匹配Footprint
	if strings.Contains(strings.ToLower(component.Package), strings.ToLower(bomItem.Footprint)) ||
		strings.Contains(strings.ToLower(bomItem.Footprint), strings.ToLower(component.Package)) {
		score += 25
	}

	// 匹配Model
	if bomItem.Model != "" && (strings.Contains(strings.ToLower(component.Model), strings.ToLower(bomItem.Model)) ||
		strings.Contains(strings.ToLower(bomItem.Model), strings.ToLower(component.Model))) {
		score += 25
	}

	// 匹配Parameters
	if bomItem.Parameters != "" && (strings.Contains(strings.ToLower(component.Name), strings.ToLower(bomItem.Parameters)) ||
		strings.Contains(strings.ToLower(component.Model), strings.ToLower(bomItem.Parameters))) {
		score += 20
	}

	return score
}

// getCellValue 获取单元格值
func getCellValue(row []string, header map[string]int, column string) string {
	if index, exists := header[column]; exists && index < len(row) {
		return strings.TrimSpace(row[index])
	}
	return ""
}

// BatchStockOut 批量出库
func (s *BOMService) BatchStockOut(batchData []struct {
	ComponentID uint   `json:"componentId"`
	Quantity    int    `json:"quantity"`
	Remark      string `json:"remark"`
}) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range batchData {
			var component models.Component
			if err := tx.First(&component, item.ComponentID).Error; err != nil {
				return err
			}

			if component.CurrentStock < item.Quantity {
				return fmt.Errorf("component %s stock insufficient: %d < %d", component.Name, component.CurrentStock, item.Quantity)
			}

			// 执行出库
			component.StockOut += item.Quantity
			component.CurrentStock -= item.Quantity

			if err := tx.Save(&component).Error; err != nil {
				return err
			}

			// 创建库存记录
			record := &models.StockRecord{
				ComponentID: item.ComponentID,
				Type:        "out",
				Quantity:    item.Quantity,
				Remark:      item.Remark,
			}
			if err := tx.Create(record).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ExportBOM 导出BOM匹配结果为Excel
func (s *BOMService) ExportBOM(bomItems []BOMItem) ([]byte, error) {
	f := excelize.NewFile()
	defer f.Close()

	// 创建工作表
	sheetName := "BOM匹配结果"
	f.NewSheet(sheetName)
	f.DeleteSheet("Sheet1")

	// 设置表头
	headers := []string{"序号", "Comment", "Footprint", "型号", "参数/规格", "数量", "匹配状态", "匹配得分", "元件名称", "库存数量", "型号", "封装"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string(rune('A'+i)))
		f.SetCellValue(sheetName, cell, header)
	}

	// 填充数据
	for i, item := range bomItems {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), i+1)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), item.Comment)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), item.Footprint)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), item.Model)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), item.Parameters)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), item.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), item.Matched)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), item.MatchScore)

		if item.Matched && item.Component != nil {
			f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), item.Component.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), item.Component.CurrentStock)
			f.SetCellValue(sheetName, fmt.Sprintf("K%d", row), item.Component.Model)
			f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), item.Component.Package)
		}
	}

	// 调整列宽
	f.SetColWidth(sheetName, "A", "L", 15)

	// 保存到字节数组
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
