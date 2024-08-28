package entities

type Sale struct {
	ID         uint
	ClientID   int
	Date       string
	Total      float64
	Company_id int
	IsLinked   bool
	DateLinked string
	Details    []SaleDetail
}

type SaleDetail struct {
	ID           uint
	SaleID       int
	ProductID    uint
	ProductClave string
	ProductName  string
	Amount       int
	Precio       float64
	Iva          float64
	Total        float64
}
