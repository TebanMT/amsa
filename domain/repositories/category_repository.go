package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type CategoryRepository interface {
	GetAllCategoriesByCompany(ctx context.Context, idCompany int) ([]entities.Category, error)
	RegisterCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
}
