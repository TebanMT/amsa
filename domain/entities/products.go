package entities

type Product struct {
	ID             uint
	Name           string
	Clave          string
	Unit           string
	Category       string
	Costo          float64
	Precio         float64
	Iva_percentage int
	Note           string
	Activated      bool
	Company_id     int
	Id_category    int
	Unit_id        int
	Type_id        int
}

type ProductType struct {
	ID          uint
	Name        string
	Description string
	Activated   bool
	Company_id  int
}
