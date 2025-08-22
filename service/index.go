package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex
// @Tags 首页
// @Accept json
// @Success 200 {string} string 你好欢迎
// @Router /index  [get]
// GetIndex 处理 /index 的 GET 请求
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "你好欢迎",
	})

}
