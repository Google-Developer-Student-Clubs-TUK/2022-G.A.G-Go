package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *subjectService) GetPosts(ctx context.Context, subjectId string, paging model.Paging) ([]model.Post, error) {
	posts, err := s.PostRepository.FindBySubjectId(ctx, subjectId, paging)
	if err != nil {
		return nil, err
	}
	fmt.Println("gorm get post success")

	return posts, err
}
