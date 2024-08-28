package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	RegisterUser(ctx context.Context, user *entities.User) error
}
