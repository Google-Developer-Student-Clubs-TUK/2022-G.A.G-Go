package model

import "context"

type UserService interface {
	TestLogin(ctx context.Context, key string, u *User) error
	DeviceRegister(ctx context.Context, uuid string) (*Device, error)
	Login(ctx context.Context, key string, u *User) error
	GetSubjects(ctx context.Context, key string, u *User, s []Subject) ([]Subject, error)
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
}

type EclassRepository interface {
	TestLogin(ctx context.Context, u *User) error
	Login(ctx context.Context, key string, u *User) error
	GetUser(ctx context.Context, u *User) error
	GetSubjects(ctx context.Context, u *User, s []Subject) ([]Subject, error)
}
