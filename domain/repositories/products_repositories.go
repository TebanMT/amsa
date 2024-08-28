package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type ProductRepository interface {
	GetAllProductsByCompanyAndType(ctx context.Context, idCompany int, idType int) ([]entities.Product, error)
	RegisterProduct(ctx context.Context, product *entities.Product) error
	UpdateProduct(ctx context.Context, product *entities.Product) error
	DeactivateProduct(ctx context.Context, idProduct int) error
	SearchProduct(ctx context.Context, nameProduct string) ([]entities.Product, error)
	GetCountProductsByType(ctx context.Context, idCompany int) (map[string]int, error)
}

type TypeProductRepository interface {
	RegisterTypeProduct(ctx context.Context, typeP *entities.ProductType) (*entities.ProductType, error)
}
