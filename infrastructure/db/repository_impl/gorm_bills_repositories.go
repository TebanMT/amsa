package repository_impl

import (
	"context"
	"time"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type BillsRepository struct {
	DB *gorm.DB
}

func NewBillRepository(db *gorm.DB) *BillsRepository {
	return &BillsRepository{DB: db}
}

func (repo *BillsRepository) RegisterBill(ctx context.Context, bill *entities.Bill) (*entities.Bill, error) {
	result := repo.DB.WithContext(ctx).Create(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return bill, nil
}

func (repo *BillsRepository) RegisterBillDetail(ctx context.Context, billDetail *entities.BillDetail) error {
	return repo.DB.WithContext(ctx).Create(billDetail).Error
}

func (repo *BillsRepository) GetAllBillsByCompany(context context.Context, companyId int) ([]entities.Bill, error) {
	var bills []entities.Bill
	if err := repo.DB.WithContext(context).Where("company_id = ?", companyId).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (repo *BillsRepository) GetAllBillDetailsByBillID(context context.Context, billId int) ([]entities.BillDetail, error) {
	var billDetails []entities.BillDetail
	if err := repo.DB.WithContext(context).Where("bill_id = ?", billId).Find(&billDetails).Error; err != nil {
		return nil, err
	}
	return billDetails, nil
}

func (repo *BillsRepository) FindBillByFolio(context context.Context, companyID int, folio string) ([]entities.Bill, error) {
	var bills []entities.Bill
	if err := repo.DB.WithContext(context).Where("company_id = ? AND folio = ?", companyID, folio).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (repo *BillsRepository) Vincular(ctx context.Context, idCompany int, id int) error {
	now := time.Now()
	formattedDate := now.Format("2006-01-02")
	return repo.DB.Model(&entities.Bill{}).Where("id = ?", id).Updates(entities.Bill{Status: "Linked", DateLinked: formattedDate}).Error
}

func (r *BillsRepository) GetAllWithClientsAndDetails(ctx context.Context, idCompany int) ([]entities.Bill, error) {
	var bills []entities.Bill
	if err := r.DB.WithContext(ctx).
		Preload("Client").
		Preload("Details").
		Where("company_id = ?", idCompany).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}
