package model

import "context"

type UserService interface {
	DeviceRegister(ctx context.Context, uuid string) (*Device, error)
	Login(ctx context.Context, key string, u *User) error
}

// repository layer
type DeviceRepository interface {
	FindByID(ctx context.Context, uuid string) (*Device, error)
	Create(ctx context.Context, d *Device) error
	Delete(ctx context.Context, uuid string) error
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
}

type EclassRepository interface {
	Login(ctx context.Context, key string, u *User) error
	GetUser(ctx context.Context, u *User) error
}
