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

func Instance() *eclass {
	if instance == nil {
		instance = &eclass{domain: "https://eclass.tukorea.ac.kr", cookies: nil}
	}
	return instance
}
