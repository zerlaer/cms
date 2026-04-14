package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 统一错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

// SuccessResponse 统一成功响应结构
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RespondWithError 发送统一格式的错误响应
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
}

// RespondWithSuccess 发送统一格式的成功响应
func RespondWithSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    data,
	})
}
