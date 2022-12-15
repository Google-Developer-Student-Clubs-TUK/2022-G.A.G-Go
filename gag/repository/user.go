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
