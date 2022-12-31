package repository

import (
	"context"
	"split_bills/config"
	"split_bills/model"
)

type userRepository struct {
	Cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository {
	return &userRepository{Cfg: cfg}
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*model.Users, error) {
	user := new(model.Users)

	if err := r.Cfg.Database().WithContext(ctx).Where("id=?", id).Preload("Bills").First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.Users) (*model.Users, error) {
	if err := r.Cfg.Database().WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepository) UpdateByID(ctx context.Context, id int, user *model.Users) (*model.Users, error) {
	if err := r.Cfg.Database().WithContext(ctx).Model(&model.Users{ID: id}).Updates(user).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepository) Delete(ctx context.Context, id int) error {
	_, err := r.FindByID(ctx, id)

	if err != nil {
		return err
	}

	res := r.Cfg.Database().WithContext(ctx).Delete(&model.Users{}, id)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *userRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Users, error) {
	var data []*model.Users

	if err := r.Cfg.Database().WithContext(ctx).Preload("Bills").Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
