package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xatta-trone/go-project-1/pkg/config"
	"github.com/xatta-trone/go-project-1/pkg/handlers"
	"github.com/xatta-trone/go-project-1/pkg/render"
)

const PORT = "localhost:3000"

func main() {
	var app config.AppConfig

	// template cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

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
