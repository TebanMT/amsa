package repository_impl

import (
	"context"
	"time"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type PurchaseOrderRepository struct {
	DB *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{DB: db}
}

func (repo *PurchaseOrderRepository) RegisterPurchaseOrder(ctx context.Context, purchase *entities.PurchaseOrder) (*entities.PurchaseOrder, error) {
	result := repo.DB.WithContext(ctx).Create(purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return purchase, nil
}

func (r *PurchaseOrderRepository) RegisterPurchaseOrderDetail(ctx context.Context, detail *entities.PurchaseOrderDetail) error {
	return r.DB.WithContext(ctx).Create(detail).Error
}

func (repo *PurchaseOrderRepository) FindPurchaseByFolio(context context.Context, companyID int, folio string) ([]entities.PurchaseOrder, error) {
	var purchase []entities.PurchaseOrder
	if err := repo.DB.WithContext(context).Where("company_id = ? AND folio = ?", companyID, folio).Find(&purchase).Error; err != nil {
		return nil, err
	}
	return purchase, nil
}

func (repo *PurchaseOrderRepository) GetAllWithSuppliersAndDetails(context context.Context, companyID int) ([]entities.PurchaseOrder, error) {
	var orders []entities.PurchaseOrder
	// Utiliza la carga conjunta para incluir proveedores y detalles de las órdenes
	if err := repo.DB.Preload("Supplier").Preload("Details").Where("company_id = ?", companyID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo *PurchaseOrderRepository) Vincular(ctx context.Context, idCompany int, id int) error {
	now := time.Now()
	formattedDate := now.Format("2006-01-02")
	return repo.DB.Model(&entities.PurchaseOrder{}).Where("id = ?", id).Updates(entities.PurchaseOrder{Status: "Linked", DateLinked: formattedDate}).Error
}
