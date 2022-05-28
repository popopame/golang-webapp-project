package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/popopame/golang-webapp-project/pkg/config"
	"github.com/popopame/golang-webapp-project/pkg/handlers"

	"github.com/popopame/golang-webapp-project/pkg/render"
)

var address = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//Change this value to true if in production
	app.InProduction = true

	//Creation of the session objet, and storage in the appconfig
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//Creation of the templatecache, and store it in the appconfig
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("err")
	}

	app.TemplateCache = templateCache
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Sprintf("Starting application on port: %s", address)

	srv := http.Server{
		Addr:    address,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
