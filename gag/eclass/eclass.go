package eclass

import (
	"net/http"

	"gag.com/eclass/model"
)

type Eclass struct {
	cookies []*http.Cookie
}

func NewEclass(e *Eclass) model.Eclass {
	return &Eclass{
		cookies: e.cookies,
	}
}
