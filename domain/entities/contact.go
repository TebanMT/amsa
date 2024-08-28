package entities

type Contact struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	ClientID uint   `json:"clientId"`
}
