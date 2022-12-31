package repository

import (
	"split_bills/config"
	"split_bills/model"
)

type billRepository struct {
	Cfg config.Config
}

func NewBillRepository(cfg config.Config) model.BillsRepository {
	return &billRepository{Cfg: cfg}
}
