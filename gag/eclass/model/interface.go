package model

import (
	"context"

	"gag.com/model"
)

type Eclass interface {
	Login(ctx context.Context, body *LoginBody) error
	GetStudent(ctx context.Context) (*Student, error)
	GetSubjects(ctx context.Context) ([]model.Subject, error)
}
