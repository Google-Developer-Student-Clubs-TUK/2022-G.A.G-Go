package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *subjectService) RegisterComment(ctx context.Context, comment *model.Comment) error {
	err := s.CommentRepository.Create(ctx, comment)
	if err != nil {
		return err
	}
	fmt.Println("gorm register comment success")

	return nil
}

func (s *subjectService) GetComments(ctx context.Context, cid uint) ([]model.Comment, error) {
	comments, err := s.CommentRepository.FindByPid(ctx, cid)
	if err != nil {
		return nil, err
	}
	fmt.Println("gorm get comment success")

	return comments, err
}

func (s *subjectService) EditComment(ctx context.Context, comment *model.Comment) error {
	err := s.CommentRepository.Update(ctx, comment)
	if err != nil {
		return err
	}
	fmt.Println("gorm edit comment success")

	return err
}

func (s *subjectService) DeleteComment(ctx context.Context, cid uint) error {
	err := s.CommentRepository.Delete(ctx, cid)
	if err != nil {
		return err
	}
	fmt.Println("gorm edit comment success")

	return err
}
