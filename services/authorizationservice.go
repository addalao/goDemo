package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type AuthorizationService struct {
	accessSecret string
}

func NewAuthorizationService(accessSecret string) *AuthorizationService {
	return &AuthorizationService{
		accessSecret: accessSecret,
	}
}

func (s *AuthorizationService) ParseTokenString(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.accessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, errors.New("parse error")
}
