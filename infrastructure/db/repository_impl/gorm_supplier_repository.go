package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{DB: db}
}

func (repo *SupplierRepository) GetAllSuppliersByCompany(ctx context.Context, idCompany int) ([]entities.Supplier, error) {
	var suppliers []entities.Supplier
	if err := repo.DB.WithContext(ctx).
		Preload("Orders").
		Preload("Orders.Details").
		Where("company_id = ?", idCompany).Find(&suppliers).Error; err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (repo *SupplierRepository) RegisterSupplier(ctx context.Context, supplier *entities.Supplier) error {
	return repo.DB.WithContext(ctx).Create(supplier).Error

}

func (repo *SupplierRepository) UpdateSupplier(context context.Context, supplier *entities.Supplier) error {
	return repo.DB.Model(&entities.Supplier{}).Where("id = ?", supplier.ID).Updates(supplier).Error
}

func (repo *SupplierRepository) DeactivateSupplier(context context.Context, idSupplier int) error {
	result := repo.DB.Model(&entities.Supplier{}).Where("id = ?", idSupplier).Update("activated", false).Error
	return result
}

func (repo *SupplierRepository) SearchSupplier(context context.Context, nameSupplier string) ([]entities.Supplier, error) {
	var suppliers []entities.Supplier
	if err := repo.DB.Where("name = ?", nameSupplier).Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (repo *SupplierRepository) GetCountSuppliers(context context.Context, idCompany int) (int64, error) {
	var count int64

	// Realiza la consulta: selecciona todos los productos, agrupa por tipo y cuenta las ocurrencias.
	err := repo.DB.WithContext(context).
		Model(&entities.Supplier{}).
		Select("COUNT(*) as count").
		Where("company_id = ? AND activated = ?", idCompany, 1).
		Count(&count).Error

	// Verifica si hubo un error en la consulta.
	if err != nil {
		return 0, err
	}

	return count, nil

}
