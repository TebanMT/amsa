package entities

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	SecondLastName string `json:"second_last_name"`
	Username       string `json:"username"`
	PhoneNumber    string `json:"phone_number"`
	CompanyName    string `json:"company_name"`
	CompanyId      string `json:"company_id"`
	Password       string `json:"password"` // El guion "-" evita que la contraseña se incluya en las respuestas JSON
}

// SetPassword encripta la contraseña y la guarda en el objeto User
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// VerifyPassword verifica si la contraseña proporcionada coincide con la contraseña del usuario.
func (u *User) VerifyPassword(password string) (bool, error) {
	fmt.Printf("PASS %s ===== %s", password, u.Password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
