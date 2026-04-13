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
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请上传文件"})
		return
	}
	defer file.Close()

	if err := h.service.Import(file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "导入成功"})
}
