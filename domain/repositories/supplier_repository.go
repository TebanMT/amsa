package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type SupplierRepository interface {
	GetAllSuppliersByCompany(ctx context.Context, idCompany int) ([]entities.Supplier, error)
	RegisterSupplier(ctx context.Context, supplier *entities.Supplier) error
	UpdateSupplier(ctx context.Context, supplier *entities.Supplier) error
	DeactivateSupplier(ctx context.Context, idSupplier int) error
	SearchSupplier(ctx context.Context, nameSupplier string) ([]entities.Supplier, error)
	GetCountSuppliers(ctx context.Context, idCompany int) (int64, error)
}
