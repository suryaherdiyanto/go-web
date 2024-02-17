package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NosurfCsrf(next http.Handler) http.Handler {
	csrf := nosurf.New(next)

	csrf.SetBaseCookie(http.Cookie{
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	return csrf
}
