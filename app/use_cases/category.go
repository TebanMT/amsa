package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type CaegoryUseCases struct {
	CategoryRepo repositories.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repositories.CategoryRepository) *CaegoryUseCases {
	return &CaegoryUseCases{
		CategoryRepo: categoryRepo,
	}
}

func (s *CaegoryUseCases) GetAllCategoriesByCompany(ctx context.Context, companyId int) ([]entities.Category, error) {
	return s.CategoryRepo.GetAllCategoriesByCompany(ctx, companyId)
}
