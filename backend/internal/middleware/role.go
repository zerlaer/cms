package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware 基于角色的权限控制中间件
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户角色
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		userRole := role.(string)
		
		// 检查用户角色是否在允许的角色列表中
		authorized := false
		for _, r := range roles {
			if strings.EqualFold(userRole, r) {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
