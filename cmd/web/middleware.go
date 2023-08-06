package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrfHanhler := nosurf.New(next)
	csrfHanhler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHanhler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
