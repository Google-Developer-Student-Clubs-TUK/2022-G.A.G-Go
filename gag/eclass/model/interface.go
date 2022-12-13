package model

import "context"

type Eclass interface {
	Login(ctx context.Context, body *LoginBody) error
	GetStudent(ctx context.Context) (*Student, error)
}
