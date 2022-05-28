package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/popopame/golang-webapp-project/pkg/config"
	"github.com/popopame/golang-webapp-project/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//init of the multiplexer
	mux := chi.NewRouter()

	//configuration of middleware
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	//definition of routes
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))


	return mux
}
