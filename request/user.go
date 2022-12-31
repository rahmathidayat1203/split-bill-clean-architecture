package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	UserRequest struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Phone   int    `json:"phone"`
		BillsID int    `json:"bills_id"`
	}
	BillsTotalRequest struct {
		ID      int `json:"id"`
		BillsID int `json:"bills_id"`
	}
)

func (req BillsTotalRequest) Validate() error {
	return validation.ValidateStruct(&req, validation.Field(&req.ID, validation.Required), validation.Field(&req.BillsID, validation.Required))
}

func (req UserRequest) Validate() error {
	return validation.ValidateStruct(&req, validation.Field(&req.Name, validation.Required), validation.Field(&req.Phone, validation.Required), validation.Field(&req.BillsID, validation.Required))
}
