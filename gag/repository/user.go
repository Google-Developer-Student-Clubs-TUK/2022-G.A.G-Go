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
		DB: db.Table("users"),
	}
}

func (r UserRepository) Create(ctx context.Context, u *model.User) error {
	r.DB.Delete(u)
	r.DB.Create(u)
	return nil
}

func (r UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{ID: id}
	r.DB.First(user)
	return user, nil
}

func (r UserRepository) SetProfileVisibility(ctx context.Context, u *model.User) error {
	u.IsProfileVisible = !u.IsProfileVisible
	r.DB.Where("id = ?", u.ID).Update("is_profile_visible", u.IsProfileVisible)
	return nil
}

func (r UserRepository) SetAlarm(ctx context.Context, u *model.User) error {
	u.IsAlarm = !u.IsAlarm
	r.DB.Where("id = ?", u.ID).Update("is_alarm", u.IsAlarm)
	return nil
}

func (r UserRepository) SetToken(ctx context.Context, u *model.User) error {
	r.DB.Where("id = ?", u.ID).Update("token", u.Token)
	return nil
}
