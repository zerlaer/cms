package service

import (
	"cms/internal/models"
	"fmt"
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ExcelService struct {
	db *gorm.DB
}

func NewExcelService(db *gorm.DB) *ExcelService {
	return &ExcelService{db: db}
}

func (s *ExcelService) Export() (*excelize.File, error) {
	var components []models.Component
	if err := s.db.Find(&components).Error; err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	sheetName := "元件清单"
	index, _ := f.NewSheet(sheetName)
	f.SetSheetName("Sheet1", sheetName)

	headers := []string{
		"商品编号", "购买渠道", "品牌", "厂家型号", "封装", "商品名称", "订购数量", "商品金额",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	for i, comp := range components {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), comp.ProductCode)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), comp.Source)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), comp.Brand)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), comp.Model)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), comp.Package)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), comp.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), comp.Quantity)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), comp.Price)
	}

	f.SetActiveSheet(index)
	return f, nil
}

func (s *ExcelService) Import(file io.Reader) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	if len(rows) < 2 {
		return nil
	}

	for i, row := range rows[1:] {
		if len(row) < 8 {
			continue
		}

		quantity, _ := strconv.Atoi(row[6])
		price, _ := strconv.ParseFloat(row[7], 64)

		source := row[1]
		// 如果没有指定购买渠道，根据商品编号自动判断
		if source == "" {
			if len(row[0]) > 0 && row[0][0] == 'C' || len(row[0]) > 0 && row[0][0] == 'c' {
				source = "LCSC"
			} else {
				source = "TB"
			}
		}

		component := &models.Component{
			ProductCode:  row[0],
			Source:       source,
			Brand:        row[2],
			Model:        row[3],
			Package:      row[4],
			Name:         row[5],
			Quantity:     quantity,
			Price:        price,
			StockIn:      quantity,
			StockOut:     0,
			CurrentStock: quantity,
		}

		if err := s.db.Create(component).Error; err != nil {
			return fmt.Errorf("导入第%d行失败：%v", i+2, err)
		}
	}

	return nil
}
