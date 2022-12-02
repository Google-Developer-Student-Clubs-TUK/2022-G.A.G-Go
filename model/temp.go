package model

type (
	TempDB struct {
		UUID       string `gorm:"primaryKey; not null"`
		PrivateKey string `gorm:"not null"`
		PublicKey  string `gorm:"not null"`
	}
)
