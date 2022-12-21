package repository

import (
	"context"

	"gag.com/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (r UserRepository) Create(ctx context.Context, u *model.User) error {
	r.DB.Create(u)
	return nil
}

func (r UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	r.DB.First(user, "id = ?", id)
	return user, nil
}

func (r UserRepository) SetProfileVisibility(ctx context.Context, u *model.User) error {
	u.IsProfileVisible = !u.IsProfileVisible
	r.DB.Table("users").Where("id = ?", u.ID).Update("is_profile_visible", u.IsProfileVisible)
	return nil
}

func (r UserRepository) SetAlarm(ctx context.Context, u *model.User) error {
	u.IsAlarm = !u.IsAlarm
	r.DB.Table("users").Where("id = ?", u.ID).Update("is_alarm", u.IsAlarm)
	return nil
}

func (r UserRepository) SetToken(ctx context.Context, u *model.User) error {
	r.DB.Table("users").Where("id = ?", u.ID).Update("token", u.Token)
	return nil
}
