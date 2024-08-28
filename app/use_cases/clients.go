package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type ClientUseCases struct {
	ClientRepo repositories.ClientRepository
}

func NewClientUseCases(repo repositories.ClientRepository) *ClientUseCases {
	return &ClientUseCases{ClientRepo: repo}
}

func (c *ClientUseCases) DeactivateClient(ctx context.Context, idUser int) (bool, error) {
	result, err := c.ClientRepo.Deactivate(ctx, idUser)
	if err != nil {
		return false, err
	}
	if result != true {
		return false, nil
	}
	return result, nil
}

func (c *ClientUseCases) UpdateClient(ctx context.Context, client *entities.Client) error {
	err := c.ClientRepo.UpdateClient(ctx, client)
	return err
}

func (s *ClientUseCases) GetCountClients(ctx context.Context, idCompany int) (int64, error) {
	ss, err := s.ClientRepo.GetCountClients(ctx, idCompany)
	if err != nil {
		return 0, err
	}
	return ss, nil
}
