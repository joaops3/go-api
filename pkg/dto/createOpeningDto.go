package dto

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

func errParamIsRequired(name, typ string) error {
	
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
	
}

type CreatingOpeningDto struct {
	Role     string `json:"role" validate:"required,min=3,max=50"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreatingOpeningDto) Validate() error {

	if err := validator.New().Struct(&r); err != nil {
		return fmt.Errorf(err.Error())
	}
	if r == nil {
		return fmt.Errorf("corpo é nil")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}



	return nil
}