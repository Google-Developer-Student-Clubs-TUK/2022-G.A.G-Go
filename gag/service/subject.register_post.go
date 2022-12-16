package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *subjectService) RegisterPost(ctx context.Context, post *model.Post) error {
	err := s.PostRepository.Create(ctx, post)
	if err != nil {
		return err
	}
	fmt.Println("gorm register post success")

	return nil
}
