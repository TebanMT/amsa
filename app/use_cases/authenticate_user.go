// app/use_cases/authenticate_user.go
package use_cases

import (
	"context"
	"fmt"

	"github.com/TebanMT/amsa/app/serializers"
	"github.com/TebanMT/amsa/domain/repositories"
	"github.com/TebanMT/amsa/domain/services"
)

type AuthenticateUserUseCase struct {
	UserRepository        repositories.UserRepository
	AuthenticationService services.AuthenticationService
}

func NewAuthenticateUser(repo repositories.UserRepository, authService services.AuthenticationService) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{UserRepository: repo, AuthenticationService: authService}
}

func (uc *AuthenticateUserUseCase) Authenticate(ctx context.Context, username, password string) (*serializers.UserDTO, error) {
	user, err := uc.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	isCorrectPassword, _ := user.VerifyPassword(password)
	if !isCorrectPassword {
		return nil, fmt.Errorf("contraseña incorrecta")
	}

	token, err := uc.AuthenticationService.GenerateToken(ctx, username)
	if err != nil {
		return nil, err
	}

	dto := serializers.NewUserDTO(user, token)

	return dto, nil
}
