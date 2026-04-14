package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"cms/internal/models"
	"cms/internal/service"
	"cms/internal/utils"
)

// 全局日志器
var logger *zap.Logger

// SetLogger 设置日志器
func SetLogger(l *zap.Logger) {
	logger = l
}

type ComponentHandler struct {
	service *service.ComponentService
}

func NewComponentHandler(service *service.ComponentService) *ComponentHandler {
	return &ComponentHandler{service: service}
}

func (h *ComponentHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	components, total, err := h.service.GetAll(page, pageSize)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     components,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (h *ComponentHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	component, err := h.service.GetByID(uint(id))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "元件不存在")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": component})
}

func (h *ComponentHandler) Create(c *gin.Context) {
	var component models.Component
	if err := c.ShouldBindJSON(&component); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(&component); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 记录创建操作日志
	if logger != nil {
		logger.Info("Component created",
			zap.Uint("id", component.ID),
			zap.String("productCode", component.ProductCode),
			zap.String("name", component.Name),
		)
	}

	utils.RespondWithSuccess(c, "创建成功", component)
}

func (h *ComponentHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var component models.Component
	if err := c.ShouldBindJSON(&component); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Update(uint(id), &component); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 记录更新操作日志
	if logger != nil {
		logger.Info("Component updated",
			zap.Uint("id", uint(id)),
			zap.String("productCode", component.ProductCode),
			zap.String("name", component.Name),
		)
	}

	utils.RespondWithSuccess(c, "更新成功", nil)
}

func (h *ComponentHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 记录删除操作日志
	if logger != nil {
		logger.Info("Component deleted",
			zap.Uint("id", uint(id)),
		)
	}

	utils.RespondWithSuccess(c, "删除成功", nil)
}

func (h *ComponentHandler) StockIn(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		Quantity int    `json:"quantity" binding:"required"`
		Remark   string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.StockIn(uint(id), req.Quantity, req.Remark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "入库成功"})
}

func (h *ComponentHandler) StockOut(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		Quantity int    `json:"quantity" binding:"required"`
		Remark   string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.StockOut(uint(id), req.Quantity, req.Remark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "出库成功"})
}

func (h *ComponentHandler) GetStockRecords(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	records, err := h.service.GetStockRecords(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}

func (h *ComponentHandler) GetAllStockRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	recordType := c.Query("type")
	componentName := c.Query("componentName")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	records, total, err := h.service.GetAllStockRecords(page, pageSize, recordType, componentName, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     records,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
