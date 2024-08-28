package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"

	"gorm.io/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{DB: db}
}

func (r *ClientRepository) CreateClient(ctx context.Context, client *entities.Client) error {
	return r.DB.WithContext(ctx).Create(client).Error
}

func (r *ClientRepository) GetAllClients(ctx context.Context, idCompany int) ([]entities.Client, error) {
	var clients []entities.Client
	if err := r.DB.WithContext(ctx).
		Preload("Contacts").
		Preload("Sales").
		Preload("Sales.Details").
		Where("company_id = ?", idCompany).Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *ClientRepository) Deactivate(ctx context.Context, idUser int) (bool, error) {
	result := r.DB.Model(&entities.Client{}).Where("id = ?", idUser).Update("activated", false)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *ClientRepository) UpdateClient(ctx context.Context, client *entities.Client) error {
	return r.DB.Model(&entities.Client{}).Where("id = ?", client.ID).Updates(client).Error
}

func (repo *ClientRepository) GetCountClients(context context.Context, idCompany int) (int64, error) {
	var results int64

	// Realiza la consulta: selecciona todos los productos, agrupa por tipo y cuenta las ocurrencias.
	err := repo.DB.WithContext(context).
		Model(&entities.Client{}).
		Select("COUNT(*) as count").
		Where("company_id = ? AND activated = ?", idCompany, 1).
		Count(&results).Error

	// Verifica si hubo un error en la consulta.
	if err != nil {
		return 0, err
	}

	return results, nil
}

func (repo *ClientRepository) GetAllClientsWithBills(ctx context.Context, idCompany int) ([]entities.Client, error) {
	var clients []entities.Client
	if err := repo.DB.WithContext(ctx).
		Preload("Bills").
		Preload("Bills.Details").
		Where("company_id = ?", idCompany).Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}
