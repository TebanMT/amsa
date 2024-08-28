package entities

type Supplier struct {
	ID          uint
	Name        string
	Executive   string
	Phone       string
	Category    string
	Address     string
	Email       string
	Description string
	Activated   bool
	Company_id  int
	Orders      []*PurchaseOrder `gorm:"foreignkey:Supplier_id"`
}
