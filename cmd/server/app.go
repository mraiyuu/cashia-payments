package server

import (
	"errors"
	"fmt"
	// "io"
	"net/http"
	"os"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mraiyuu/cashia-payments/internal/routes"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.Heartbeat("/"))

	port := ":8000"

	fmt.Printf("server running on port %s\n", port)

	//register routes
	routes.RegisterAuthRoutes(r)

	err := http.ListenAndServe(port, r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		os.Exit(1)
	}
}
