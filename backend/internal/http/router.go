package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/mirily/shorty/internal/http/handlers"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", handlers.Health)

	return router
}
