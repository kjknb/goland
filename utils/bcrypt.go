/**
* @Auth:ShenZ
* @Description: Bcrypt密码加密工具类
* @CreateDate:2022/06/15 16:27:35
 */
package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// BcryptCost 设置bcrypt的计算成本(默认值: 10)
var BcryptCost = 10

// GenerateBcryptPassword 使用Bcrypt加密密码
func GenerateBcryptPassword(plainpwd string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainpwd), BcryptCost)
	if err != nil {
		return "", fmt.Errorf("密码加密失败: %v", err)
	}
	return string(hashedBytes), nil
}

// CompareBcryptPassword 验证Bcrypt加密的密码
func CompareBcryptPassword(plainpwd, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainpwd))
	return err == nil
}

// SetBcryptCost 设置Bcrypt的计算成本
func SetBcryptCost(cost int) {
	if cost < 4 {
		cost = 4
	} else if cost > 31 {
		cost = 31
	}
	BcryptCost = cost
}

// 为了保持与原有代码的兼容性

// MakePassword 兼容原有方法的Bcrypt实现
// 注意：Bcrypt不需要盐值，salt参数将被忽略
func MakePassword(plainpwd, salt string) string {
	hashedPassword, err := GenerateBcryptPassword(plainpwd)
	if err != nil {
		return ""
	}
	return hashedPassword
}

// ValidPassword 兼容原有方法的Bcrypt实现
// 注意：Bcrypt不需要盐值，salt参数将被忽略
func ValidPassword(plainpwd, salt string, password string) bool {
	return CompareBcryptPassword(plainpwd, password)
}
