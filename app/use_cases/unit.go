package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type MeasurementUnitUseCases struct {
	UnitRepo repositories.MeasurementUnitRepository
}

func NewUnitUseCase(unitRepo repositories.MeasurementUnitRepository) *MeasurementUnitUseCases {
	return &MeasurementUnitUseCases{
		UnitRepo: unitRepo,
	}
}

func (s *MeasurementUnitUseCases) GetAllUnitsByCompany(ctx context.Context, companyId int) ([]entities.MeasurementUnit, error) {
	return s.UnitRepo.GetAllUnitsByCompany(ctx, companyId)
}
