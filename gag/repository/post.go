package repository

import (
	"context"

	"gag.com/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) model.PostRepository {
	return PostRepository{
		DB: db,
	}
}

func (r PostRepository) Create(ctx context.Context, p *model.Post) error {
	r.DB.Create(p)
	return nil
}

func (r PostRepository) FindBySubjectId(ctx context.Context, subjectId string, paging model.Paging) ([]model.Post, error) {
	posts := make([]model.Post, 0)
	r.DB.Where("sid = ?", subjectId).Find(posts)

	offset := paging.Page * paging.PerPage

	if err := r.DB.
		Where("sid = ?", subjectId).
		Order("created_at DESC").
		Offset(offset).
		Limit(paging.PerPage).
		Find(posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r PostRepository) Update(ctx context.Context, p *model.Post) error {
	nowPost := &model.Post{}
	r.DB.First(nowPost, "id = ?", p.ID)
	r.DB.Model(nowPost).Updates(p)
	return nil
}

func (r PostRepository) Delete(ctx context.Context, pid uint) error {
	r.DB.Delete(pid)
	return nil
}
