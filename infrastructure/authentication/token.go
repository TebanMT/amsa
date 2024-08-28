package authentication

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Importar otros paquetes necesarios

type JWTTokenService struct {
	secretKey string
}

func NewJWTTokenService(secretKey string) *JWTTokenService {
	return &JWTTokenService{secretKey: secretKey}
}

func (generator *JWTTokenService) GenerateToken(ctx context.Context, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(generator.secretKey))
}

func (s *JWTTokenService) VerifyToken(ctx context.Context, tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}
