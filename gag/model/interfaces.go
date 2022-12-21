package model

import "context"

type UserService interface {
	TestLogin(ctx context.Context, key string, u *User) error
	DeviceRegister(ctx context.Context, uuid string) (*Device, error)
	Login(ctx context.Context, key string, u *User) error
	GetSubjects(ctx context.Context, key string, u *User, s []Subject) ([]Subject, error)
	GetTodos(ctx context.Context, key string, u *User) ([]Todo, error)
	SetProfileVisibility(ctx context.Context, key string, u *User) error
	SetAlarm(ctx context.Context, key string, u *User) error
	SetToken(ctx context.Context, key string, u *User) error
	GetProfile(ctx context.Context, key string, u *User) error
}

type SubjectService interface {
	GetPosts(ctx context.Context, subjectId string, pagination Pagination) ([]Post, error)
	RegisterPost(ctx context.Context, post *Post) error
	EditPost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, pid uint) error
	GetComments(ctx context.Context, pid uint) ([]Comment, error)
	RegisterComment(ctx context.Context, comment *Comment) error
	EditComment(ctx context.Context, comment *Comment) error
	DeleteComment(ctx context.Context, cid uint) error
}

// repository layer
type DeviceRepository interface {
	FindByID(ctx context.Context, uuid string) (*Device, error)
	Create(ctx context.Context, d *Device) error
	Delete(ctx context.Context, uuid string) error
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	SetProfileVisibility(ctx context.Context, u *User) error
	SetAlarm(ctx context.Context, u *User) error
	SetToken(ctx context.Context, u *User) error
}

type PostRepository interface {
	Create(ctx context.Context, p *Post) error
	FindBySubjectId(ctx context.Context, subjectId string, pagination Pagination) ([]Post, error)
	Update(ctx context.Context, p *Post) error
	Delete(ctx context.Context, pid uint) error
}

type CommentRepository interface {
	Create(ctx context.Context, p *Comment) error
	FindByPid(ctx context.Context, pid uint) ([]Comment, error)
	Update(ctx context.Context, p *Comment) error
	Delete(ctx context.Context, pid uint) error
}

type EclassRepository interface {
	TestLogin(ctx context.Context, u *User) error
	Login(ctx context.Context, key string, u *User) error
	GetUser(ctx context.Context, u *User) error
	GetSubjects(ctx context.Context, u *User, s []Subject) ([]Subject, error)
	GetTodos(ctx context.Context, u *User) ([]Todo, error)
}
