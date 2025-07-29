package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("sua-chave-secreta-super-segura")

// Claims personalizadas (opcional)
type CustomClaims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, email, password string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Email:    email,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "auth_microservice",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}
		return secretKey, nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, fmt.Errorf("malformed token")
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, fmt.Errorf("expired token")
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, fmt.Errorf("invalid token")
		case errors.Is(err, jwt.ErrSignatureInvalid):
			return nil, fmt.Errorf("signature invalid")
		default:
			return nil, err
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.UserID == "" || claims.Email == "" {
			return nil, fmt.Errorf("invalid claims")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
