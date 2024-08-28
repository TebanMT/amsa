package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type GetAllClientsUseCase struct {
	ClientRepo repositories.ClientRepository
}

func NewGetAllClientsUseCase(clientRepo repositories.ClientRepository) *GetAllClientsUseCase {
	return &GetAllClientsUseCase{
		ClientRepo: clientRepo,
	}
}

func (uc *GetAllClientsUseCase) Execute(ctx context.Context, idCompany int) ([]entities.Client, error) {
	return uc.ClientRepo.GetAllClients(ctx, idCompany)
}

func (s *GetAllClientsUseCase) GetAllClientsWithBillsAndDetailsByCompany(ctx context.Context, companyID int) ([]entities.Client, error) {
	orders, err := s.ClientRepo.GetAllClientsWithBills(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
