package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xatta-trone/go-project-1/pkg/config"
	"github.com/xatta-trone/go-project-1/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	chi := chi.NewRouter()
	chi.Use(middleware.Recoverer)
	// chi.Use(WriteToConsole)
	chi.Use(NoSurf)
	chi.Use(SessionLoad)

	chi.Get("/", handlers.Repo.Home)
	chi.Get("/about", handlers.Repo.About)

	return chi
}
