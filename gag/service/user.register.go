package service

import (
	"context"

	"gag.com/model"
	"gag.com/util"
)

func (s *userService) DeviceRegister(ctx context.Context, uuid string) (*model.Device, error) {
	rh := util.RSAHelper{}
	rh.GenerateKey(1024)

	private_key, _ := rh.PrivateToStringPEM()
	public_key, _ := rh.PublicToStringPEM()

	device := &model.Device{
		UUID:          uuid,
		RsaPrivateKey: private_key,
		RsaPublicKey:  public_key,
	}
	err := s.DeviceRepository.Create(ctx, device)

	return device, err
}
