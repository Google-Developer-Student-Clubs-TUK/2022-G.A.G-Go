package model

type (
	Paging struct {
		Page    int `json:"page"`
		PerPage int `json:"per_page"`
	}
)
