package tool

import (
	"github.com/golang-jwt/jwt/v4"
	"goDemo/internal/config"
	"strconv"
	"time"
)

// GenerateToken 生成令牌
func GenerateToken(userID uint, config config.Config) (string, error) {

	id := strconv.FormatUint(uint64(userID), 10)
	// 创建一个声明（Claims）
	claims := jwt.MapClaims{
		"userId": id,
		"exp":    time.Now().Add(time.Minute * time.Duration(config.Auth.AccessExpire)).Unix(), // 令牌过期时间（例如，1天后）
	}
	// 使用秘钥签名令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Auth.AccessSecret)) //

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
