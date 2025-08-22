package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex 处理 /index 的 GET 请求
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}
