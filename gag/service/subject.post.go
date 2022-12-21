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

func (s *subjectService) GetPosts(ctx context.Context, subjectId string, pagination model.Pagination) ([]model.Post, error) {
	posts, err := s.PostRepository.FindBySubjectId(ctx, subjectId, pagination)
	if err != nil {
		return nil, err
	}
	fmt.Println("gorm get post success")

	return posts, err
}

func (s *subjectService) EditPost(ctx context.Context, post *model.Post) error {
	err := s.PostRepository.Update(ctx, post)
	if err != nil {
		return err
	}
	fmt.Println("gorm edit post success")

	return err
}

func (s *subjectService) DeletePost(ctx context.Context, pid uint) error {
	err := s.PostRepository.Delete(ctx, pid)
	if err != nil {
		return err
	}
	fmt.Println("gorm edit post success")

	return err
}
