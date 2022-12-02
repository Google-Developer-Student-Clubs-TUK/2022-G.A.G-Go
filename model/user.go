package model

import (
	"time"
)

type (
	UserDB struct {
		StudentId  string `gorm:"primaryKey; not null"`
		UUID       string `gorm:"not null"`
		PrivateKey string `gorm:"not null"`
		AES        string `gorm:"not null"`
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
