package schemas

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID int 
	Email string
	Password string
}


func NewUser(id int, email string, password string) *User {
	return &User{
		Email: email, Password: password, ID: id,
	}
}