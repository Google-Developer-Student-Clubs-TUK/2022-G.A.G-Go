package model

import (
	"time"
)

type (
	User struct {
		// 보안 계정 정보
		ID            string `gorm:"primaryKey; not null" json:"id"`
		UUID          string `gorm:"not null" json:"uuid"`
		RsaPrivateKey string `gorm:"not null" json:"rsa_private_key"`
		AesPassword   string `gorm:"not null" json:"aes_password"`

		// 개인 정보
		Name     string `gorm:"not null" json:"name"`
		Email    string `gorm:"not null" json:"email"`
		ImageURL string `gorm:"not null" json:"image_url"`
		Cookie   string `gorm:"not null" json:"cookie"`

		// timestamp
		CreatedAt time.Time `gorm:"not null" json:"created_at"`
		UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	}
)
