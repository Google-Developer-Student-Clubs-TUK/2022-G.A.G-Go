package service

import (
	"context"
	"fmt"

	"gag.com/model"
)

func (s *userService) TestLogin(ctx context.Context, key string, u *model.User) error {
	_, err := s.DeviceRepository.FindByID(ctx, u.UUID)
	if err != nil {
		return err
	}

	// ------- id가 암호화 되어 있지 않다면 불필요 코드 -----
	// RSA 복호화

	// AES 복호화

	// eclass login
	err = s.EclassRepository.TestLogin(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println("eclass login success")

	// eclass get user
	err = s.EclassRepository.GetUser(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println("eclass get user success")
	// 로그인 성공시 DB 저장
	err = s.UserRepository.Create(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println("user create success")

	// 디바이스 정보 삭제
	err = s.DeviceRepository.Delete(ctx, u.UUID)
	if err != nil {
		return err
	}
	fmt.Println("device delete success")

	return err
}
