package model

type Subject struct {
	name      string `json:"name"`
	professor string `json:"professor"`
	room      string `json:"room"`
	time      string `json:"time"`
}
