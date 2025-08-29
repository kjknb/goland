package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT 密钥（生产环境应从配置中读取，不要硬编码）
var jwtSecret = []byte("your-secret-key") // 请修改为更安全的密钥

// Claims 结构体
type Claims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(identity, name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // token 有效期为 24 小时

	claims := Claims{
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "GInchat",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析 JWT token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("解析token失败: %v", err)
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
		return nil, fmt.Errorf("token无效")
	}

	return nil, err
}

// JWTAuthMiddleware JWT 认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// 提取实际的 token (去掉 "Bearer " 前缀)
		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "token格式错误",
			})
			c.Abort()
			return
		}

		token := authHeader[len(bearerPrefix):]

		// 解析 token
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "token无效或已过期: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 将解析出的信息存储到上下文
		c.Set("identity", claims.Identity)
		c.Set("name", claims.Name)
		c.Next()
	}
}
