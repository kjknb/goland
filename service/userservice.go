package service

import (
	"GInchat/models"
	"GInchat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
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
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册",
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "密码不一致",
		})
		return
	}
	hashedPassword := utils.MakePassword(password, "")
	user.Password = hashedPassword
	fmt.Printf("原始密码: %s\n", password)
	fmt.Printf("Bcrypt加密后: %s\n", hashedPassword)

	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})
}

// FindUserByNameAndPwd
// @Summary 所有用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/GetUserList [get]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name = c.Query("name")
	password := c.Query("password")
	data

	c.JSON(200, gin.H{
		"message": data,
	})

}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
		"data":    user,
	})

}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update:", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "修改参数不匹配！",
			"data":    user,
		})

	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0, //  0成功   -1失败
			"message": "修改用户成功！",
			"data":    user,
		})
	}

}
