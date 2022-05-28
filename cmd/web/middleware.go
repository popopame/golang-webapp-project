package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//A middleware take a http.handler as input, and a http.handler as output, so it take input from request and output to another handler

// GenerateCSRFToken add protection to all post type request 
func GenerateCSRFToken(next http.Handler) http.Handler{

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad Loads and save session on every request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}