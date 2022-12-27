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
		DB: db.Table("devices"),
	}
}

func (r DeviceRepository) Create(ctx context.Context, d *model.Device) error {
	r.DB.Delete(d)
	r.DB.Create(d)
	return nil
}

func (r DeviceRepository) FindByID(ctx context.Context, uuid string) (*model.Device, error) {
	device := &model.Device{UUID: uuid}
	r.DB.First(device)
	return device, nil
}

func (r DeviceRepository) Delete(ctx context.Context, uuid string) error {
	r.DB.Delete(&model.Device{})
	return nil
}
