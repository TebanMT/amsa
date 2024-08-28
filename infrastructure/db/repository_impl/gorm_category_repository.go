package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (repo *CategoryRepository) GetAllCategoriesByCompany(ctx context.Context, idCompany int) ([]entities.Category, error) {
	var categories []entities.Category
	if err := repo.DB.WithContext(ctx).Where("company_id = ?", idCompany).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repo *CategoryRepository) RegisterCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	result := repo.DB.WithContext(ctx).Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
