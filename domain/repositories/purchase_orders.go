package repositories

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
)

type PurchaseOrderRepository interface {
	RegisterPurchaseOrder(ctx context.Context, purchase *entities.PurchaseOrder) (*entities.PurchaseOrder, error)
	RegisterPurchaseOrderDetail(ctx context.Context, purchaseDetail *entities.PurchaseOrderDetail) error
	FindPurchaseByFolio(ctx context.Context, companyID int, folio string) ([]entities.PurchaseOrder, error)
	GetAllWithSuppliersAndDetails(ctx context.Context, companyID int) ([]entities.PurchaseOrder, error)
	Vincular(ctx context.Context, idCompany int, id int) error
}
