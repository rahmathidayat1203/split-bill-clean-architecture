package request

import validation "github.com/go-ozzo/ozzo-validation"

type (
	BillsRequest struct {
		BillsName          string `json:"bills_name"`
		ManyPeopleForBills int    `json:"many_people_for_bills"`
	}
)

func (req BillsRequest) Validate() error {
	return validation.ValidateStruct(&req, validation.Field(&req.BillsName, validation.Required), validation.Field(&req.ManyPeopleForBills, validation.Required))
}
