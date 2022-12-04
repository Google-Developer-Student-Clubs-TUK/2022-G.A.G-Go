package model

import (
	"time"
)

type (
	UserDB struct {
		Id         string `gorm:"primaryKey; not null"`
		UUID       string `gorm:"not null"`
		PrivateKey string `gorm:"not null"`
		AesData    string `gorm:"not null"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
