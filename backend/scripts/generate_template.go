package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	sheetName := "元件导入模板"
	index, _ := f.NewSheet(sheetName)
	f.SetSheetName("Sheet1", sheetName)

	headers := []string{
		"商品编号", "购买渠道", "品牌", "厂家型号", "封装", "商品名称", "订购数量", "商品金额",
	}

	style, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 12},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#CCCCCC"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
		f.SetCellStyle(sheetName, cell, cell, style)
	}

	examples := [][]interface{}{
		{"RES-001", "TB", "Vishay", "CRCW080510K0FKEA", "0805", "贴片电阻", "1000", "0.01"},
		{"CAP-001", "LCSC", "Murata", "GRM188R71H104KA93D", "0603", "贴片电容", "5000", "0.05"},
		{"IC-001", "TB", "TI", "LM358DR", "SOP-8", "运算放大器", "500", "1.20"},
	}

	for i, row := range examples {
		rowNum := i + 2
		for j, value := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, rowNum)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	f.SetColWidth(sheetName, "A", "H", 15)
	f.SetActiveSheet(index)

	templateDir := "template"
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		os.MkdirAll(templateDir, 0755)
	}

	templatePath := filepath.Join(templateDir, "元件导入模板.xlsx")
	if err := f.SaveAs(templatePath); err != nil {
		fmt.Println("错误:", err)
		return
	}

	fmt.Println("✅ 模板文件已生成:", templatePath)
	fmt.Println("\n📋 导入模板说明:")
	fmt.Println("- 商品编号：商品的产品编号")
	fmt.Println("- 购买渠道：TB 或 LCSC")
	fmt.Println("- 品牌：元件品牌")
	fmt.Println("- 厂家型号：制造商定义的型号")
	fmt.Println("- 封装：元件封装类型")
	fmt.Println("- 商品名称：元件名称")
	fmt.Println("- 订购数量：采购数量")
	fmt.Println("- 商品金额：单价（元）")
	fmt.Println("\n⚠️  注意：导入时请保留表头，从第 2 行开始填写数据")
	fmt.Println("💡 系统会自动生成 ID，无需手动填写序号")
}