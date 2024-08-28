package use_cases

import (
	"context"
	"fmt"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type PurchaseOrderUseCase struct {
	PurchaseOrderRepo repositories.PurchaseOrderRepository
	SupplierRepo      repositories.SupplierRepository
}

type PurchaseOrderUseCaseOption func(*PurchaseOrderUseCase)

func WithSupplierRepo(supplierRepo repositories.SupplierRepository) PurchaseOrderUseCaseOption {
	return func(puc *PurchaseOrderUseCase) {
		puc.SupplierRepo = supplierRepo
	}
}

func NewPurchaseOrderUseCase(purchaseRepo repositories.PurchaseOrderRepository, opts ...PurchaseOrderUseCaseOption) *PurchaseOrderUseCase {
	puc := &PurchaseOrderUseCase{
		PurchaseOrderRepo: purchaseRepo,
	}
	for _, opt := range opts {
		opt(puc)
	}
	return puc
}

func (s *PurchaseOrderUseCase) RegisterPurchaseOrder(ctx context.Context, purchase *entities.PurchaseOrder, purchaseDetails []*entities.PurchaseOrderDetail) error {
	registredPurchase, err := s.PurchaseOrderRepo.RegisterPurchaseOrder(ctx, purchase)
	if err != nil {
		return err
	}
	// Registrar cada detalle de venta
	for _, detail := range purchaseDetails {
		detail.OrderID = int(registredPurchase.ID)
		if err := s.PurchaseOrderRepo.RegisterPurchaseOrderDetail(ctx, detail); err != nil {
			return err
		}
	}
	return nil
}

func (s *PurchaseOrderUseCase) ExistFolio(ctx context.Context, companyID int, folio string) (bool, error) {
	bills, err := s.PurchaseOrderRepo.FindPurchaseByFolio(ctx, companyID, folio)
	fmt.Println("RESULT use case: ", bills)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false, err
	}
	if len(bills) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (s *PurchaseOrderUseCase) GetOrdersWithDetailsAndSupliers(ctx context.Context, companyID int) ([]entities.PurchaseOrder, error) {
	orders, err := s.PurchaseOrderRepo.GetAllWithSuppliersAndDetails(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *PurchaseOrderUseCase) Vincular(ctx context.Context, idCompany int, id int) (bool, error) {

	err := s.PurchaseOrderRepo.Vincular(ctx, idCompany, id)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false, err
	}
	return true, nil

}
