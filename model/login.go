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

type Device struct {
	// json tag to de-serialize json body
	UUID string `json:"name"`
}

type RSA struct {
	// json tag to de-serialize json body
	PublicKey string `json:"name"`
}
