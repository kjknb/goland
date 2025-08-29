package service

import (
	"GInchat/models"
	"GInchat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList
// @Summary 获取所有用户列表
// @Description 获取系统中所有用户的列表信息
// @Tags 用户模块
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/list [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取用户列表成功",
		"data":    data,
		"count":   len(data),
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Param repassword formData string true "确认密码"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	// 从表单获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")

	// 验证用户名是否已存在
	if data := models.FindUserByName(name); data.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "用户名已注册",
		})
		return
	}

	// 验证密码一致性
	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "密码不一致",
		})
		return
	}

	// 密码加密处理
	hashedPassword := utils.MakePassword(password)
	fmt.Printf("原始密码: %s\n", password)
	fmt.Printf("Bcrypt加密后: %s\n", hashedPassword)

	// 创建用户
	user := models.UserBasic{
		Name:     name,
		Password: hashedPassword,
	}
	models.CreateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "新增用户成功",
	})
}

// FindUserByNameAndPwd
// @Summary 所有用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	// 参数验证
	if name == "" || password == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名和密码不能为空",
		})
		return
	}

	// 查找用户
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    404,
			"message": "用户不存在",
		})
		return
	}

	// 验证密码
	flag := utils.ValidPassword(password, user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":    401,
			"message": "密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.Identity, user.Name)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "token生成失败",
		})
		return
	}

	// 返回用户信息和token
	userResponse := gin.H{
		"identity": user.Identity,
		"name":     user.Name,
		"token":    token,
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "登录成功",
		"data":    userResponse,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	// 获取并验证ID参数
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数id不能为空",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数id格式错误",
		})
		return
	}

	// 检查用户是否存在
	user := models.FindUserByID(uint(id))
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}

	// 执行删除操作
	models.DeleteUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除用户成功",
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 用户模块
// @Accept multipart/form-data
// @Produce json
// @Param id formData int true "用户ID"
// @Param name formData string false "用户名"
// @Param password formData string false "密码"
// @Param phone formData string false "电话"
// @Param email formData string false "邮箱"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/updateUser [put]
func UpdateUser(c *gin.Context) {
	// 获取并验证ID参数
	idStr := c.PostForm("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数id不能为空",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "参数id格式错误",
		})
		return
	}

	// 检查用户是否存在
	existingUser := models.FindUserByID(uint(id))
	if existingUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "用户不存在",
		})
		return
	}

	// 准备更新数据
	user := models.UserBasic{}
	user.ID = uint(id)
	user.Name = c.PostForm("name")

	// 处理密码更新（如果需要）
	if password := c.PostForm("password"); password != "" {
		user.Password = utils.MakePassword(password)
	} else {
		// 如果不更新密码，保持原密码
		user.Password = existingUser.Password
	}

	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	fmt.Println("update:", user)

	// 验证数据
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "修改参数不匹配！",
			"error":   err.Error(),
		})
		return
	}

	// 执行更新
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "修改用户成功！",
		"data":    user,
	})
}
