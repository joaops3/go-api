package dto

import "fmt"

type UpdateOpeningDto struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningDto) Validate() error {
	if r == nil {
		return fmt.Errorf("corpo é nil")
	}
	if r.Role != "" ||  r.Company != "" || r.Location != "" || r.Link != "" || r.Salary > 0 {
		return nil
	}
	return  fmt.Errorf("corpo é vaziu, um campo deve ser enviado")
}