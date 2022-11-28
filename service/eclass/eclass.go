package eclass

import (
	"net/http"
)

type eclass struct {
	domain  string
	cookies []*http.Cookie
}

// singleton
var instance *eclass

func Instance(cookies []*http.Cookie) *eclass {
	if instance == nil {
		instance = &eclass{domain: "https://eclass.tukorea.ac.kr", cookies: cookies}
	}
	return instance
}
