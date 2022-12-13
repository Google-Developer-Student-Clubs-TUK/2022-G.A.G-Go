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
