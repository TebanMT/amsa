package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type TypeProductRepository struct {
	DB *gorm.DB
}

func NewTypeProductRepository(db *gorm.DB) *TypeProductRepository {
	return &TypeProductRepository{DB: db}
}

func (repo *TypeProductRepository) RegisterTypeProduct(ctx context.Context, typeP *entities.ProductType) (*entities.ProductType, error) {
	result := repo.DB.WithContext(ctx).Table("product_type").Create(typeP)
	if result.Error != nil {
		return nil, result.Error
	}
	return typeP, nil
}
