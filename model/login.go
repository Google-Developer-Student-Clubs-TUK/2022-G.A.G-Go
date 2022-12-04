package model

type (
	Login struct {
		UUID    string `json:"uuid"`
		Key     string `json:"key"`
		AesData string `json:"login_data"`
	}

	LoginBody struct {
		Usr_id  string `json:"usr_id"`
		Usr_pwd string `json:"usr_pwd"`
	}
)
