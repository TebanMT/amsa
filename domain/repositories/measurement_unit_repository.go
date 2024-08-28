package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type MeasurementUnitRepository interface {
	GetAllUnitsByCompany(ctx context.Context, idCompany int) ([]entities.MeasurementUnit, error)
	RegisterUnit(ctx context.Context, category *entities.MeasurementUnit) (*entities.MeasurementUnit, error)
}
