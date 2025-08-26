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
	Phone         string
	Email         string
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

func CreateUser(user UserBasic) *gorm.DB {

	return utils.DB.Create(&user)

}
