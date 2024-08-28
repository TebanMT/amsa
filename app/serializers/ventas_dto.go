package serializers

import "github.com/TebanMT/amsa/domain/entities"

type VentaDetailsDTO struct {
	ID           uint    `json:"id"`
	SaleID       int     `json:"sale_id"`
	ProductClave string  `json:"clave"`
	ProductoName string  `json:"producto_name"`
	ProductoID   uint    `json:"producto_id"`
	Cantidad     int     `json:"cantidad"`
	De           string  `json:"de"`
	Precio       float64 `json:"precio"`
	Iva          float64 `json:"iva"`
	Total        float64 `json:"total"`
}

type VentaDTO struct {
	ID         uint `json:"id"`
	ClientID   int  `json:"client_id"`
	Ventas     []VentaDetailsDTO
	CompanyId  int    `json:"company_id"`
	IsLinked   bool   `json:"is_linked"`
	DateLinked string `json:"date_linked"`
}

func NewVentaDTO(sale *entities.Sale, ventas []VentaDetailsDTO) *VentaDTO {
	return &VentaDTO{
		ID:         sale.ID,
		ClientID:   sale.ClientID,
		Ventas:     ventas,
		CompanyId:  sale.Company_id,
		IsLinked:   sale.IsLinked,
		DateLinked: sale.DateLinked,
	}
}

func NewVentaDetailsDTO(detail *entities.SaleDetail) *VentaDetailsDTO {
	return &VentaDetailsDTO{
		ID:           detail.ID,     // Suponiendo que ID de producto está disponible
		SaleID:       detail.SaleID, // Suponiendo que ID de producto está disponible
		ProductClave: detail.ProductClave,
		ProductoName: detail.ProductName,
		ProductoID:   detail.ProductID, // Asumiendo que esto es un identificador único del producto
		Cantidad:     detail.Amount,
		De:           "stock",
		Precio:       detail.Precio,
		Iva:          detail.Iva,
		Total:        detail.Total,
	}
}
