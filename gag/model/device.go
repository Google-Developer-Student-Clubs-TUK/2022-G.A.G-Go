package model

type (
	Device struct {
		// 디바이스 ID
		UUID string `gorm:"primaryKey; not null" json:"uuid"`

		// RSA Key
		RsaPrivateKey string `gorm:"not null" json:"rsa_private_key"`
		RsaPublicKey  string `gorm:"not null" json:"rsa_public_key"`
	}
)
