package repository

import (
	"context"
	"split_bills/config"
	"split_bills/model"
)

type billRepository struct {
	Cfg config.Config
}

func NewBillRepository(cfg config.Config) model.BillsRepository {
	return &billRepository{Cfg: cfg}
}
func (r *billRepository) FindByID(ctx context.Context, id int) (*model.Bills, error) {
	bill := new(model.Bills)

	if err := r.Cfg.Database().WithContext(ctx).Where("id=?", id).First(bill).Error; err != nil {
		return nil, err
	}
	return bill, nil
}

func (r *billRepository) Create(ctx context.Context, bill *model.Bills) (*model.Bills, error) {
	if err := r.Cfg.Database().WithContext(ctx).Create(&bill).Error; err != nil {
		return nil, err
	}
	return bill, nil
}

func (r *billRepository) UpdateByID(ctx context.Context, id int, bill *model.Bills) (*model.Bills, error) {
	if err := r.Cfg.Database().WithContext(ctx).Model(&model.Bills{ID: id}).Updates(bill).Find(bill).Error; err != nil {
		return nil, err
	}
	return bill, nil
}
func (r *billRepository) Delete(ctx context.Context, id int) error {
	_, err := r.FindByID(ctx, id)

	if err != nil {
		return err
	}

	res := r.Cfg.Database().WithContext(ctx).Delete(&model.Bills{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *billRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Bills, error) {
	var data []*model.Bills

	if err := r.Cfg.Database().WithContext(ctx).Select("id", "bills_name", "many_people_for_bills", "created_at", "updated_at").Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
