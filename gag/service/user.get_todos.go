package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *userService) GetTodos(ctx context.Context, key string, u *model.User) ([]model.Todo, error) {
	user, err := s.UserRepository.FindByID(ctx, u.ID)
	todos := []model.Todo{}
	if err != nil {
		return todos, err
	}
	fmt.Println("user find success")
	err = s.EclassRepository.Login(ctx, key, user)
	if err != nil {
		return todos, err
	}
	fmt.Println("eclass login success")

	// eclass get user
	todos, err = s.EclassRepository.GetTodos(ctx, user)
	if err != nil {
		return todos, err
	}
	fmt.Println("eclass get user success")

	return todos, err
}
