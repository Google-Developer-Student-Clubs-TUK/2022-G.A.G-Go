package model

type (
	Subject struct {
		ID        string `json:"id" gorm:"primaryKey"`
		Name      string `json:"name"`
		StartTime string `json:"startTime"`
		IsPinned  bool   `json:"isPinned"`
		Room      string `json:"room"`
	}
)
