package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// Валидация токена
func ValidateToken(tokenString string) (*jwt.StandardClaims, error) {
	// Парсим токен
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи верный
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем валидность токена
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil // Возвращаем claims вместо *jwt.Token
	}

	return nil, errors.New("invalid token")
}
