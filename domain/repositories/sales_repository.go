package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type SalesRepository interface {
	RegisterSale(ctx context.Context, sale *entities.Sale) (*entities.Sale, error)
	RegisterSaleDetail(ctx context.Context, saleDetail *entities.SaleDetail) error
	GetAllSalesByCompany(ctx context.Context, idCompany int) ([]entities.Sale, error)
	GetAllSaleDetailsBySaleID(ctx context.Context, saleID int) ([]entities.SaleDetail, error)
	Vincularr(ctx context.Context, idCompany int, id int) error
}
