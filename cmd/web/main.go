package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/xatta-trone/go-project-1/pkg/config"
	"github.com/xatta-trone/go-project-1/pkg/handlers"
	"github.com/xatta-trone/go-project-1/pkg/render"
)

const PORT = "localhost:3000"
var app config.AppConfig
var sessionManager *scs.SessionManager


func main() {

	// change to true in production 
	app.InProduction = false

	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction



	// template cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.Session = sessionManager

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
