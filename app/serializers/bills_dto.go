package serializers

import "github.com/TebanMT/amsa/domain/entities"

type BillDetailsDTO struct {
	ID           uint    `json:"id"`
	BillID       int     `json:"bill_id"`
	ProductClave string  `json:"clave"`
	ProductoName string  `json:"producto_name"`
	ProductoID   uint    `json:"producto_id"`
	Cantidad     int     `json:"cantidad"`
	De           string  `json:"de"`
	Precio       float64 `json:"precio"`
	Iva          float64 `json:"iva"`
	Total        float64 `json:"total"`
}

type BillDTO struct {
	ID        uint   `json:"id"`
	Folio     string `json:"folio"`
	ClientID  int    `json:"client_id"`
	Status    string `json:"status"`
	Facturas  []BillDetailsDTO
	CompanyId int     `json:"company_id"`
	Total     float64 `json:"total"`
	//IsLinked   bool    `json:"is_linked"`
	//DateLinked string  `json:"date_linked"`
}

func NewBillDTO(bill *entities.Bill, bills []BillDetailsDTO) *BillDTO {
	return &BillDTO{
		ID:        bill.ID,
		Folio:     bill.Folio,
		ClientID:  bill.Client_id,
		Status:    bill.Status,
		Total:     bill.Total,
		Facturas:  bills,
		CompanyId: bill.Company_id,
		//IsLinked:   bill.IsLinked,
		//DateLinked: bill.DateLinked,
	}
}

func NewBillDetailsDTO(detail *entities.BillDetail) *BillDetailsDTO {
	return &BillDetailsDTO{
		ID:           detail.ID,     // Suponiendo que ID de producto está disponible
		BillID:       detail.BillID, // Suponiendo que ID de producto está disponible
		ProductClave: detail.ProductClave,
		ProductoName: detail.ProductName,
		ProductoID:   detail.ProductID, // Asumiendo que esto es un identificador único del producto
		Cantidad:     detail.Quantity,
		De:           "almacen",
		Precio:       detail.UnitaryPrice,
		Iva:          detail.Iva,
		Total:        detail.TotalPrice,
	}
}
