package api

import (
	"log"
	"main/api/fetch"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Address string
	Version string
	Db      DbConfig
}

type DbConfig struct {
	Address            string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}

type Application struct {
	Config Config
	Store  fetch.Storage
}

func (api *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	//set a timeout value on the request context (ctx), that will signal
	//through ctx.Done() that the request hast timed out and further processing should be stopped
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		// Swagger UI
		r.Get("/docs", api.SwaggerHandler)
		r.Get("/docs/swagger.json", api.SwaggerJSONHandler)

		// Health check
		r.Get("/health", api.HealthCheckHandler)

		// Deck routes
		r.Route("/decks", func(r chi.Router) {
			r.Get("/", api.GetDecksHandler)
			r.Post("/", api.CreateDeckHandler)
			r.Get("/featured", api.GetFeatureDecksHandler)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", api.GetDeckHandler)
				r.Put("/", api.UpdateDeckHandler)
				r.Delete("/", api.DeleteDeckHandler)
				r.Get("/cards", api.GetDeckCardsHandler)
			})
		})

		// Activity routes
		r.Get("/activity", api.GetRecentActivityHandler)
	})

	return r
}

func (api *Application) Run(handler http.Handler) error {

	srv := &http.Server{
		Addr:         api.Config.Address,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	serverAddress := "http://localhost" + api.Config.Address + "/" + api.Config.Version + "/health"
	log.Printf("Starting server at %s", serverAddress)

	return srv.ListenAndServe()
}
