package model

type (
	Todo struct {
		ID       string `json:"subjectId"`
		Name     string `json:"name"`
		DeadLine string `json:"deadLine"`
		IsDone   bool   `json:"isDone"`
	}
)
