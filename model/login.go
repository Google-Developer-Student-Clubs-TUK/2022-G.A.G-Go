package model

type (
	Login struct {
		Id  string `json:"id"`
		Pwd string `json:"pwd"`
	}

	EclassLoginBody struct {
		Usr_id  string `json:"usr_id"`
		Usr_pwd string `json:"usr_pwd"`
	}
)
