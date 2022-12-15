package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *userService) GetSubjects(ctx context.Context, key string, u *model.User, sl []model.Subject) ([]model.Subject, error) {
	user, err := s.UserRepository.FindByID(ctx, u.ID)
	if err != nil {
		return sl, err
	}
	fmt.Println("user find success")
	err = s.EclassRepository.TestLogin(ctx, user)
	if err != nil {
		return sl, err
	}
	fmt.Println("eclass login success")

	// eclass get user
	sl, err = s.EclassRepository.GetSubjects(ctx, user, sl)
	if err != nil {
		return sl, err
	}
	fmt.Println("eclass get user success")

	return sl, err
}
