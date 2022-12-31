package model

import (
	"context"
	"time"
)

const (
	TransactionsBillStatusFinished = "finished"
	TransactionsBillStatusDebit    = "debit"
)

type (
	Transactions struct {
		ID          int       `json:"id"`
		BillsName   string    `json:"bills_name"`
		BillsID     int       `json:"bills_id"`
		Bills       *Bills    `json:"bills"`
		DetailBills []string  `json:"details_bills"`
		BillsStatus string    `json:"bills_status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	TransactionsRepository interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transactions, error)
	}
	TransactionsUsecase interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transactions, int, error)
	}
)
