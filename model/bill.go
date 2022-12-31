package model

import (
	"context"
	"split_bills/request"
	"time"
)

type (
	Bills struct {
		ID                 int       `json:"id"`
		BillsName          string    `json:"bills_name"`
		ManyPeopleForBills int       `json:"many_people_for_bills"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	}
	BillsRepository interface {
		Create(ctx context.Context, bills *Bills) (*Bills, error)
		UpdateByID(ctx context.Context, id int, bills *Bills) (*Bills, error)
		FindByID(ctx context.Context, id int) (*Bills, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*Bills, error)
	}
	BillsUseCase interface {
		GetByID(ctx context.Context, id int) (*Bills, error)
		FetchBills(ctx context.Context, limit, offset int) ([]*Bills, error)
		DestroyBills(ctx context.Context, id int) error
		EditBills(ctx context.Context, id int, req *request.BillsRequest) (*Bills, error)
		StoreBills(ctx context.Context, req *request.BillsRequest) (*Bills, error)
	}
)
