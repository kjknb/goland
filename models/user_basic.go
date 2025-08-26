package models

import (
	"GInchat/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClentIp       string
	ClentTime     string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LogOutTime    *time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"

}
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data

}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(name string) *gorm.DB {
	Phone := UserBasic{}
	return utils.DB.Where("Phone = ?", Phone).First(&Phone)

}
func FindUserByEmail(name string) *gorm.DB {
	Email := UserBasic{}
	return utils.DB.Where("Email = ?", Email).First(&Email)

}

func CreateUser(user UserBasic) *gorm.DB {

	return utils.DB.Create(&user)

}

func DeleteUser(user UserBasic) *gorm.DB {

	return utils.DB.Delete(&user)

}
func UpdateUser(user UserBasic) *gorm.DB {

	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone, Email: user.Email})

}
