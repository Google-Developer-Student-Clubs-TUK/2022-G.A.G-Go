package service

import (
	"context"
	"fmt"

	"gag.com/model"
	"github.com/ulule/deepcopier"
)

func (s *userService) GetProfile(ctx context.Context, key string, u *model.User) error {
	var err error
	user, err := s.UserRepository.FindByID(ctx, u.ID)
	if err != nil {
		return err
	}
	fmt.Println("eclass login success")
	deepcopier.Copy(user).To(u)
	return err
}
