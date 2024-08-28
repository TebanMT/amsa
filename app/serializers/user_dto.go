package serializers

import "github.com/TebanMT/amsa/domain/entities"

type UserDTO struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Token          string `json:"token"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	SecondLastName string `json:"second_last_name"`
	CompanyName    string `json:"company_name"`
	CompanyId      string `json:"company_id"`
}

func NewUserDTO(user *entities.User, token string) *UserDTO {
	return &UserDTO{
		ID:             user.ID,
		Username:       user.Username,
		Name:           user.Name,
		LastName:       user.LastName,
		SecondLastName: user.SecondLastName,
		Token:          token,
		CompanyName:    user.CompanyName,
		CompanyId:      user.CompanyId,
	}
}
