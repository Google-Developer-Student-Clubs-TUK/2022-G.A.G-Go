package model

import "gorm.io/gorm"

type (
	Post struct {
		gorm.Model
		Writer string `gorm:"not null" json:"Writer"`
		SID    string `gorm:"not null" json:"sid"`

		// 게시물 정보
		Title   string `gorm:"not null" json:"title"`
		Content string `gorm:"not null" json:"content"`
	}
)
