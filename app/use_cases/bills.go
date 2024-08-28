package use_cases

import (
	"context"
	"fmt"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type BillsUseCase struct {
	BillRepo repositories.BillsRepository
	SaleRepo repositories.SalesRepository
}

type BillUseCaseOption func(*BillsUseCase)

func WithBillRepoAndSaleRepo(saleRepo repositories.SalesRepository) BillUseCaseOption {
	return func(puc *BillsUseCase) {
		puc.SaleRepo = saleRepo
	}
}

func NewBillUseCase(billRepo repositories.BillsRepository, opts ...BillUseCaseOption) *BillsUseCase {
	puc := &BillsUseCase{
		BillRepo: billRepo,
	}
	for _, opt := range opts {
		opt(puc)
	}
	return puc
}

func (s *BillsUseCase) RegisterBill(ctx context.Context, bill *entities.Bill, billDetails []*entities.BillDetail) error {
	sal, err := s.BillRepo.RegisterBill(ctx, bill)
	if err != nil {
		return err
	}
	// Registrar cada detalle de venta
	for _, detail := range billDetails {
		detail.BillID = int(sal.ID) // Asigna el ID de la venta a cada detalle
		if err := s.BillRepo.RegisterBillDetail(ctx, detail); err != nil {
			return err // Podrías manejar los errores para intentar guardar tantos detalles como sea posible antes de fallar
		}
	}
	return nil
}

func (s *BillsUseCase) GetAllBillsAndDetailsByCompany(ctx context.Context, companyID int) ([]entities.Bill, error) {
	orders, err := s.BillRepo.GetAllWithClientsAndDetails(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *BillsUseCase) ExistFolio(ctx context.Context, companyID int, folio string) (bool, error) {
	bills, err := s.BillRepo.FindBillByFolio(ctx, companyID, folio)
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

func (s *BillsUseCase) Vincular(ctx context.Context, idCompany int, id int) (bool, error) {

	err := s.BillRepo.Vincular(ctx, idCompany, id)
	if err != nil {
		fmt.Printf("%s\n", err)
		return false, err
	}
	return true, nil

}
