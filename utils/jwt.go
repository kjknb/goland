package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getJwtSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// 开发环境默认密钥，生产环境务必设置环境变量
		return []byte("dev_secret_key_please_change_in_production")
	}
	return []byte(secret)
}

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
func GenerateToken(username string, userID uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // token有效期为24小时

	claims := Claims{
		username,
		userID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "GInchat",
			IssuedAt:  nowTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(getJwtSecret()) // 修改这里

	return token, err
}

// ParseToken 解析JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJwtSecret(), nil // 修改这里
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// RefreshToken 刷新token
func RefreshToken(token string) (string, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	return GenerateToken(claims.Username, claims.UserID)
}
