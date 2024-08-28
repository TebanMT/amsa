package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type CreateClientUseCase struct {
	ClientRepo repositories.ClientRepository
}

func NewCreateClientUseCase(clientRepo repositories.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientRepo: clientRepo,
	}
}

func (uc *CreateClientUseCase) Execute(ctx context.Context, client *entities.Client) error {
	return uc.ClientRepo.CreateClient(ctx, client)
}
