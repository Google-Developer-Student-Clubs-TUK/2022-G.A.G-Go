package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *userService) SetProfileVisibility(ctx context.Context, key string, u *model.User) error {
	err := s.UserRepository.SetProfileVisibility(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println("eclass get user success")

	return err
}
