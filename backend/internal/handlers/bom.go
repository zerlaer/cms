package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"cms/internal/service"
	"cms/internal/utils"
)

type BOMHandler struct {
	service *service.BOMService
}

func NewBOMHandler(service *service.BOMService) *BOMHandler {
	return &BOMHandler{service: service}
}

// ImportBOM 导入BOM文件
func (h *BOMHandler) ImportBOM(c *gin.Context) {
	// 限制文件大小为10MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)

	file, err := c.FormFile("file")
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "No file uploaded")
		return
	}

	// 保存文件
	ext := filepath.Ext(file.Filename)
	if ext != ".xlsx" && ext != ".xls" {
		utils.RespondWithError(c, http.StatusBadRequest, "Only Excel files are allowed")
		return
	}

	// 打开上传的文件进行魔数验证
	src, err := file.Open()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to open uploaded file")
		return
	}
	defer src.Close()

	// 读取文件头部进行魔数验证
	buffer := make([]byte, 8)
	_, err = src.Read(buffer)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to read file header")
		return
	}

	// 验证Excel文件魔数
	// xls: D0 CF 11 E0 A1 B1 1A E1
	// xlsx: 50 4B 03 04
	if !((buffer[0] == 0xD0 && buffer[1] == 0xCF && buffer[2] == 0x11 && buffer[3] == 0xE0 && buffer[4] == 0xA1 && buffer[5] == 0xB1 && buffer[6] == 0x1A && buffer[7] == 0xE1) ||
		(buffer[0] == 0x50 && buffer[1] == 0x4B && buffer[2] == 0x03 && buffer[3] == 0x04)) {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid Excel file")
		return
	}

	// 创建临时目录
	tempDir := "./temp"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create temp directory")
		return
	}

	filePath := filepath.Join(tempDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to save file")
		return
	}
	defer os.Remove(filePath)

	// 导入BOM
	bomItems, err := h.service.ImportBOM(filePath)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to import BOM: "+err.Error())
		return
	}

	// 匹配库存
	matchedItems, err := h.service.MatchBOM(bomItems)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to match BOM: "+err.Error())
		return
	}

	// 记录操作日志
	if logger != nil {
		logger.Info("BOM imported and matched",
			zap.Int("totalItems", len(matchedItems)),
			zap.Int("matchedItems", countMatchedItems(matchedItems)),
		)
	}

	utils.RespondWithSuccess(c, "BOM imported and matched successfully", matchedItems)
}

// BatchStockOut 批量出库
func (h *BOMHandler) BatchStockOut(c *gin.Context) {
	var batchData []struct {
		ComponentID uint   `json:"componentId"`
		Quantity    int    `json:"quantity"`
		Remark      string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&batchData); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data: "+err.Error())
		return
	}

	// 验证数据
	for _, item := range batchData {
		if item.ComponentID == 0 {
			utils.RespondWithError(c, http.StatusBadRequest, "Component ID is required")
			return
		}
		if item.Quantity <= 0 {
			utils.RespondWithError(c, http.StatusBadRequest, "Quantity must be greater than 0")
			return
		}
	}

	if err := h.service.BatchStockOut(batchData); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to batch stock out: "+err.Error())
		return
	}

	// 记录操作日志
	if logger != nil {
		logger.Info("Batch stock out completed",
			zap.Int("items", len(batchData)),
		)
	}

	utils.RespondWithSuccess(c, "Batch stock out completed successfully", nil)
}

// countMatchedItems 统计匹配的项目数量
func countMatchedItems(items []service.BOMItem) int {
	count := 0
	for _, item := range items {
		if item.Matched {
			count++
		}
	}
	return count
}

// ExportBOM 导出BOM匹配结果
func (h *BOMHandler) ExportBOM(c *gin.Context) {
	var bomItems []service.BOMItem

	if err := c.ShouldBindJSON(&bomItems); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data: "+err.Error())
		return
	}

	// 导出Excel
	data, err := h.service.ExportBOM(bomItems)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to export BOM: "+err.Error())
		return
	}

	// 设置响应头
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=BOM匹配结果.xlsx")
	c.Header("Content-Length", string(len(data)))

	// 写入响应
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}
