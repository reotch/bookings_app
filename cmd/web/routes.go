package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/reotch/bookings_app/pkg/config"
	"github.com/reotch/bookings_app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// router, aka: "multiplexer", abbr. "mux"
	mux := chi.NewRouter()

	// Middleware comes after the router declaration
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// GET routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
