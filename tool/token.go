package tool

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"goDemo/internal/config"
	"time"
)

type Claims struct {
	user_id int
	exp     int64
	jwt.RegisteredClaims
}

// GenerateToken 生成令牌
func GenerateToken(userID uint, config config.Config) (string, error) {
	// 创建一个声明（Claims）
	claims := jwt.MapClaims{
		"user_id": int(userID),
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 令牌过期时间（例如，1天后）
	}
	// 使用秘钥签名令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SecretKey)) //  config.SecretKey = zhwsecret777
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string, config config.Config) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("noToken")
}
