package entities

type Client struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone,omitempty"`
	Address     string    `json:"address,omitempty"`
	Email       string    `json:"email,omitempty"`
	Description string    `json:"description,omitempty"`
	Activated   bool      `json:"activated,omitempty"`
	Company_id  int       `json:"company_id"`
	Contacts    []Contact `json:"contacts,omitempty"`
	Sales       []Sale    `json:"sales,omitempty"`
	Bills       []*Bill   `gorm:"foreignkey:Client_id"`
}
