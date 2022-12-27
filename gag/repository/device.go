package repository

import (
	"context"

	"gag.com/model"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	DB *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) model.DeviceRepository {
	return DeviceRepository{
		DB: db,
	}
}

func (r DeviceRepository) Create(ctx context.Context, d *model.Device) error {
	r.DB.Delete(d.UUID)
	r.DB.Create(d)
	return nil
}

func (r DeviceRepository) FindByID(ctx context.Context, uuid string) (*model.Device, error) {
	device := &model.Device{}
	r.DB.First(device, "uuid = ?", uuid)
	return device, nil
}

func (r DeviceRepository) Delete(ctx context.Context, uuid string) error {
	r.DB.Delete(uuid)
	return nil
}
