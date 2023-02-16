package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xatta-trone/go-project-1/pkg/config"
	"github.com/xatta-trone/go-project-1/pkg/handlers"
	"github.com/xatta-trone/go-project-1/pkg/render"
)

const PORT = "3000"

func main() {
	var app config.AppConfig

	// template cache 
	tc,err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplate(&app)



	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/divide", handlers.Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT))

	http.ListenAndServe(fmt.Sprintf("localhost:%s", PORT), nil)
}
