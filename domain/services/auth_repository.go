package services

import (
	"context"
)

type AuthenticationService interface {
	GenerateToken(ctx context.Context, username string) (string, error)
	VerifyToken(ctx context.Context, token string) (bool, error)
}
