package v1

import (
	"bez/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func NewRouter(router *chi.Mux, googleAPI usecase.GoogleAPI) {
	newAuthRoutes(router, googleAPI)
}
