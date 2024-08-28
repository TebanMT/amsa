package repository_impl

import (
	"context"
	"time"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type SalesRepository struct {
	DB *gorm.DB
}

func NewSaleRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{DB: db}
}

func (repo *SalesRepository) RegisterSale(ctx context.Context, sale *entities.Sale) (*entities.Sale, error) {
	result := repo.DB.WithContext(ctx).Create(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return sale, nil
}

func (repo *SalesRepository) RegisterSaleDetail(ctx context.Context, saleDetail *entities.SaleDetail) error {
	return repo.DB.WithContext(ctx).Create(saleDetail).Error
}

func (repo *SalesRepository) GetAllSalesByCompany(context context.Context, companyId int) ([]entities.Sale, error) {
	var sales []entities.Sale
	if err := repo.DB.WithContext(context).Where("company_id = ?", companyId).Find(&sales).Error; err != nil {
		return nil, err
	}
	return sales, nil
}

func (repo *SalesRepository) GetAllSaleDetailsBySaleID(context context.Context, saleId int) ([]entities.SaleDetail, error) {
	var saleDetails []entities.SaleDetail
	if err := repo.DB.WithContext(context).Where("sale_id = ?", saleId).Find(&saleDetails).Error; err != nil {
		return nil, err
	}
	return saleDetails, nil
}

func (repo *SalesRepository) Vincularr(ctx context.Context, idCompany int, id int) error {
	now := time.Now()
	formattedDate := now.Format("2006-01-02")
	return repo.DB.Model(&entities.Sale{}).Where("id = ?", id).Updates(entities.Sale{IsLinked: true, DateLinked: formattedDate}).Error
}
