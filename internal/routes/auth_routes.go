package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mraiyuu/cashia-payments/internal/handlers"
)

func RegisterAuthRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Post("/initiateAuth", handlers.AuthenticateMerchant)
	})
}
