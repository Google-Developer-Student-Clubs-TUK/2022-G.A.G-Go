package model

type (
	Login struct {
		UUID      string `json:"uuid"`
		PublicKey string `json:"public_key"`
		AES       string `json:"aes"`
		IV        string `json:"iv"`
	}

	LoginBody struct {
		Usr_id  string `json:"usr_id"`
		Usr_pwd string `json:"usr_pwd"`
	}
)
