package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type BillsRepository interface {
	RegisterBill(ctx context.Context, bill *entities.Bill) (*entities.Bill, error)
	RegisterBillDetail(ctx context.Context, billDetail *entities.BillDetail) error
	GetAllBillsByCompany(ctx context.Context, idCompany int) ([]entities.Bill, error)
	GetAllBillDetailsByBillID(ctx context.Context, billID int) ([]entities.BillDetail, error)
	FindBillByFolio(ctx context.Context, companyID int, folio string) ([]entities.Bill, error)
	Vincular(ctx context.Context, idCompany int, id int) error
	GetAllWithClientsAndDetails(ctx context.Context, idCompany int) ([]entities.Bill, error)
}
