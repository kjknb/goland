package service

import (
	"GInchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 首页
// @Success 200 {string} JSON{"code","message"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})

}

// CreateUser
// @Tags 首页
// @Success 200 {string} JSON{"code","message"}
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	Password := c.Query("password")
	rePassword := c.Query("repassord")
	if Password != rePassword {
		return
	}
	c.JSON(-1, gin.H{
		"message": "密码不一致",
	})
	user.Password = Password
	models.CreateUser(user)
	c.JSON(100, gin.H{
		"message": "新增用户成功",
	})

}
