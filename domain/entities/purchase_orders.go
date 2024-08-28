package entities

type PurchaseOrder struct {
	ID          uint
	Folio       string
	Supplier_id int
	OrderDate   string
	Total       float64
	Company_id  int
	Status      string
	Supplier    *Supplier             `gorm:"foreignKey:Supplier_id"`
	Details     []PurchaseOrderDetail `gorm:"foreignkey:OrderID"`
	DateLinked  string
}

type PurchaseOrderDetail struct {
	ID           uint
	OrderID      int
	ProductID    uint
	ProductClave string
	ProductName  string
	Quantity     int
	UnitaryPrice float64
	TotalPrice   float64
	Iva          float64
}
