package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/mirily/shorty/internal/http/handlers"
)

func NewRouter(
	authHandler *handlers.AuthHandler,
) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", handlers.Health)

	router.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
		})
	})

	return router
}
