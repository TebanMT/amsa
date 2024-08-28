package use_cases

import (
	"context"

	"github.com/TebanMT/amsa/app/serializers"
	"github.com/TebanMT/amsa/domain/entities"
	"github.com/TebanMT/amsa/domain/repositories"
)

type SaleUseCase struct {
	SaleRepo repositories.SalesRepository
}

func NewSaleUseCase(saleRepo repositories.SalesRepository) *SaleUseCase {
	return &SaleUseCase{
		SaleRepo: saleRepo,
	}
}

func (s *SaleUseCase) RegisterSale(ctx context.Context, sale *entities.Sale, saleDetails []*entities.SaleDetail) error {
	sal, err := s.SaleRepo.RegisterSale(ctx, sale)
	if err != nil {
		return err
	}
	// Registrar cada detalle de venta
	for _, detail := range saleDetails {
		detail.SaleID = int(sal.ID) // Asigna el ID de la venta a cada detalle
		if err := s.SaleRepo.RegisterSaleDetail(ctx, detail); err != nil {
			return err // Podrías manejar los errores para intentar guardar tantos detalles como sea posible antes de fallar
		}
	}
	return nil
}

func (s *SaleUseCase) GetAllSalesAndDetailsByCompany(ctx context.Context, companyId int) ([]serializers.VentaDTO, error) {
	sales, err := s.SaleRepo.GetAllSalesByCompany(ctx, companyId)
	if err != nil {
		return nil, err // Podrías manejar los errores para intentar guardar tantos detalles como sea posible antes de fallar
	}

	//dto := serializers.NewVentaDTO(user, token)
	salesDTO := make([]serializers.VentaDTO, 0, len(sales))
	for _, sale := range sales {
		details, err := s.SaleRepo.GetAllSaleDetailsBySaleID(ctx, int(sale.ID))
		if err != nil {
			return nil, err // Podrías manejar los errores para intentar guardar tantos detalles como sea posible antes de fallar
		}

		detailsDTO := make([]serializers.VentaDetailsDTO, 0, len(details))
		for _, detail := range details {
			dtoDetails := serializers.NewVentaDetailsDTO(&detail)
			detailsDTO = append(detailsDTO, *dtoDetails)
		}
		dtoSale := serializers.NewVentaDTO(&sale, detailsDTO)

		salesDTO = append(salesDTO, *dtoSale)
	}
	return salesDTO, nil
}
