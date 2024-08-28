package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type SupplierUseCases struct {
	SupplierRepo repositories.SupplierRepository
}

func NewSupplierUseCase(supplierRepo repositories.SupplierRepository) *SupplierUseCases {
	return &SupplierUseCases{
		SupplierRepo: supplierRepo,
	}
}

func (s *SupplierUseCases) GetAllSuppliersByCompany(ctx context.Context, companyId int) ([]entities.Supplier, error) {
	return s.SupplierRepo.GetAllSuppliersByCompany(ctx, companyId)
}

func (s *SupplierUseCases) RegisterSupplierIntoDB(ctx context.Context, supplier *entities.Supplier) error {
	return s.SupplierRepo.RegisterSupplier(ctx, supplier)
}

func (s *SupplierUseCases) UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error {
	return s.SupplierRepo.UpdateSupplier(ctx, supplier)
}

func (s *SupplierUseCases) DeactivateSupplier(ctx context.Context, idSupplier int) error {
	return s.SupplierRepo.DeactivateSupplier(ctx, idSupplier)
}

func (s *SupplierUseCases) SearchSupplier(ctx context.Context, nameSupplier string) ([]entities.Supplier, error) {
	return s.SupplierRepo.SearchSupplier(ctx, nameSupplier)
}

func (s *SupplierUseCases) GetCountSuppliers(ctx context.Context, idCompany int) (int64, error) {
	ss, err := s.SupplierRepo.GetCountSuppliers(ctx, idCompany)
	if err != nil {
		return 0, err
	}
	return ss, nil
}
