package model

import "gorm.io/gorm"

type (
	Comment struct {
		// 개인 정보
		gorm.Model
		Writer string `gorm:"not null" json:"writer"`
		PID    string `gorm:"not null" json:"pid"`

		// 댓글 내용
		Content string `gorm:"not null" json:"content"`
		Index   uint   `gorm:"autoIncrement; not null" json:"index"`
	}
)
