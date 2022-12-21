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
		DB: db,
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

func (r CommentRepository) FindByCid(ctx context.Context, cid string) (*model.Comment, error) {
	comment := &model.Comment{}
	r.DB.First(comment, "id = ?", cid)
	return comment, nil
}

func (r CommentRepository) Update(ctx context.Context, c *model.Comment) error {
	nowComment := &model.Comment{}
	r.DB.First(nowComment, "id = ?", c.ID)
	r.DB.Model(nowComment).Updates(c)
	return nil
}

func (r CommentRepository) Delete(ctx context.Context, cid uint) error {
	r.DB.Delete(cid)
	return nil
}
