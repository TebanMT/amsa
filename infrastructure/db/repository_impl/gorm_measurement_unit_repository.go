package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type MeasurementUnitRepository struct {
	DB *gorm.DB
}

func NewUnitRepository(db *gorm.DB) *MeasurementUnitRepository {
	return &MeasurementUnitRepository{DB: db}
}

func (repo *MeasurementUnitRepository) GetAllUnitsByCompany(ctx context.Context, idCompany int) ([]entities.MeasurementUnit, error) {
	var unit []entities.MeasurementUnit
	if err := repo.DB.WithContext(ctx).Where("company_id = ?", idCompany).Find(&unit).Error; err != nil {
		return nil, err
	}
	return unit, nil
}

func (repo *MeasurementUnitRepository) RegisterUnit(ctx context.Context, category *entities.MeasurementUnit) (*entities.MeasurementUnit, error) {
	result := repo.DB.WithContext(ctx).Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
