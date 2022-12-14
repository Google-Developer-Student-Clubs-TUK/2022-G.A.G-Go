package repository

import (
	"context"

	"gag.com/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) model.CommentRepository {
	return CommentRepository{
		DB: db.Table("comments"),
	}
}

func (r CommentRepository) Create(ctx context.Context, c *model.Comment) error {
	r.DB.Create(c)
	return nil
}

func (r CommentRepository) FindByPid(ctx context.Context, pid uint) ([]model.Comment, error) {
	comments := make([]model.Comment, 0)
	r.DB.Where("pid = ?", pid).Find(comments)

	return comments, nil
}

func (r CommentRepository) FindByCid(ctx context.Context, cid uint) (*model.Comment, error) {
	comment := &model.Comment{}
	comment.ID = cid
	r.DB.First(comment)
	return comment, nil
}

func (r CommentRepository) Update(ctx context.Context, c *model.Comment) error {
	r.DB.Where(c).Updates(c)
	return nil
}

func (r CommentRepository) Delete(ctx context.Context, cid uint) error {
	comment := &model.Comment{}
	comment.ID = cid
	r.DB.Delete(comment)
	return nil
}
