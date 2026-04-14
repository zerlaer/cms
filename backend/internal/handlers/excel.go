package handlers

import (
	"cms/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExcelHandler struct {
	service *service.ExcelService
}

func NewExcelHandler(service *service.ExcelService) *ExcelHandler {
	return &ExcelHandler{service: service}
}

func (h *ExcelHandler) Export(c *gin.Context) {
	f, err := h.service.Export()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=components.xlsx")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *ExcelHandler) Import(c *gin.Context) {
	// 限制文件大小为10MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20) // 10MB
	
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传文件"})
		return
	}
	defer file.Close()
	
	// 验证文件类型
	ext := header.Filename[len(header.Filename)-4:]
	if ext != ".xlsx" && ext != ".xls" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "只支持Excel文件(.xlsx, .xls)"})
		return
	}

	if err := h.service.Import(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "导入成功"})
}
