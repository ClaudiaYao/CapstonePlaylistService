package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *PlaylistService) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	// mux.Route("/playlists", func(r chi.Router) {
	// 	r.Get("/", app.Playlists)
	// 	r.Post("/new", app.CreatePlaylist)

	// 	r.Get("/{code}", h.internalPlan.Get)
	// 	r.Put("/{code}", h.internalPlan.Update)
	// })

	mux.Post("/playlist/new", app.CreatePlaylist)
	mux.Get("/", app.Welcome)

	mux.Get("/playlists", app.Playlists)
	// mux.Get("/playlists/sort?{}", app.Playlists)
	return mux
}
