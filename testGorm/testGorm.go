package main

import (
	"GInchat/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	defer closeDB(db)

	// 创建用户
	userID := createUser(db, "申专")
	if userID == 0 {
		log.Fatal("创建用户失败")
	}

	// 查询用户
	user, err := getUserByID(db, userID)
	if err != nil {
		log.Fatal("查询用户失败:", err)
	}
	fmt.Printf("查询到的用户: %+v\n", user)

	// 更新用户密码
	err = updateUserPassword(db, userID, "1234")
	if err != nil {
		log.Fatal("更新密码失败:", err)
	}
	fmt.Println("密码更新成功")

	// 验证更新结果
	updatedUser, _ := getUserByID(db, userID)
	fmt.Printf("更新后的用户: %+v\n", updatedUser)
}

// 初始化数据库连接
func initDB() *gorm.DB {
	dsn := "root:123mdx0.0@tcp(127.0.0.1:3306)/GinChat1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	return db
}

// 关闭数据库连接
func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("获取数据库实例失败:", err)
		return
	}
	sqlDB.Close()
}

// 创建用户
func createUser(db *gorm.DB, name string) uint {
	user := &models.UserBasic{Name: name}
	result := db.Create(user)
	if result.Error != nil {
		log.Println("创建用户失败:", result.Error)
		return 0
	}
	fmt.Printf("创建用户成功，ID: %d\n", user.ID)
	return user.ID
}

// 根据ID查询用户
func getUserByID(db *gorm.DB, id uint) (*models.UserBasic, error) {
	var user models.UserBasic
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// 更新用户密码
func updateUserPassword(db *gorm.DB, id uint, password string) error {
	result := db.Model(&models.UserBasic{}).Where("id = ?", id).Update("password", password)
	return result.Error
}
