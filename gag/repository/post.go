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
		DB: db.Table("posts"),
	}
}

func (r PostRepository) Create(ctx context.Context, p *model.Post) error {
	r.DB.Create(p)
	return nil
}

func (r PostRepository) FindBySubjectId(ctx context.Context, subjectId string, pagination model.Pagination) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	offset := pagination.Page * pagination.PerPage

	if err := r.DB.
		Where("sid = ?", subjectId).
		Order("created_at DESC").
		Offset(offset).
		Limit(pagination.PerPage).
		Find(posts).Error; err != nil {
		return posts, nil
	}

	return posts, nil
}

func (r PostRepository) FindByPostId(ctx context.Context, pid string) (*model.Post, error) {
	post := &model.Post{}
	r.DB.First(post, "id = ?", pid)
	return post, nil
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
