package model

import (
	"context"
	"split_bills/request"
	"time"
)

type (
	Users struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Phone     int       `json:"phone"`
		BillsID   int       `json:"bills_id"`
		Bills     *Bills    `json:"bills"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	UserRepository interface {
		Create(ctx context.Context, user *Users) (*Users, error)
		UpdateByID(ctx context.Context, id int, user *Users) (*Users, error)
		FindByID(ctx context.Context, id int) (*Users, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*Users, error)
	}
	UserUsecase interface {
		GetByID(ctx context.Context, id int) (*Users, error)
		FetchUser(ctx context.Context, limit, offset int) ([]*Users, error)
		DestroyUser(ctx context.Context, id int) error
		EditUser(ctx context.Context, id int, req *request.UserRequest) (*Users, error)
		StoreUser(ctx context.Context, req *request.UserRequest) (*Users, error)
		BillsTotal(ctx context.Context, req *request.BillsTotalRequest) error
	}
)
