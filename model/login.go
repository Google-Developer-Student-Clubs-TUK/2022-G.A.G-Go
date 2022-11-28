package model

type (
	Login struct {
		Id  string `json:"id"`
		Pwd string `json:"pwd"`
	}

	EclassLoginBody struct {
		Usr_id  string
		Usr_pwd string
	}
)
