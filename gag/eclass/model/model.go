package model

type (
	LoginBody struct {
		Usr_id  string `json:"usr_id"`
		Usr_pwd string `json:"usr_pwd"`
	}
)

type (
	Student struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		ImageUrl string `json:"imageUrl"`
	}
)
