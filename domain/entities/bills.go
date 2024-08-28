package entities

type Bill struct {
	ID         uint
	Folio      string
	Client_id  int
	Date       string
	Total      float64
	Company_id int
	Status     string
	Client     *Client `gorm:"foreignkey:Client_id"`
	Details    []BillDetail
	//IsLinked   bool
	DateLinked string
}

type BillDetail struct {
	ID           uint
	BillID       int
	ProductID    uint
	ProductClave string
	ProductName  string
	Quantity     int
	UnitaryPrice float64
	Iva          float64
	TotalPrice   float64
}
