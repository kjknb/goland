package middleware

import (
	"GInchat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "请提供认证token",
			})
			c.Abort()
			return
		}

		// 检查token格式，通常为"Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token格式错误，应为Bearer <token>",
			})
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的token",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		// 将用户信息存储在context中，以便后续使用
		c.Set("username", claims.Username)
		c.Set("user_claims", claims)

		c.Next()
	}
}

// OptionalAuthMiddleware 可选的认证中间件，不强制要求认证
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// 检查token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.Next()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.Next()
			return
		}

		// 将用户信息存储在context中
		c.Set("username", claims.Username)
		c.Set("user_claims", claims)

		c.Next()
	}
}
