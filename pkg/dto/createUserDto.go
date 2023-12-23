package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateUserDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}

func (u *CreateUserDto) Validate() error {
	if err := validator.New().Struct(&u); err != nil {
		return fmt.Errorf(err.Error())
	} 
	return nil
}	