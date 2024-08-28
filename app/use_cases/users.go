package use_cases

import (
	"context"
	"fmt"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type UserUseCases struct {
	UserRepository repositories.UserRepository
}

func NewUserCases(repo repositories.UserRepository) *UserUseCases {
	return &UserUseCases{UserRepository: repo}
}

func (uc *UserUseCases) RegisterUserIntoDB(ctx context.Context, usert *entities.User) error {
	err := usert.SetPassword(usert.Password)
	fmt.Printf("USER == %s", usert)
	if err != nil {
		// Manejar el error
	}
	errr := uc.UserRepository.RegisterUser(ctx, usert)
	if errr != nil {
		return err
	}

	return nil
}
