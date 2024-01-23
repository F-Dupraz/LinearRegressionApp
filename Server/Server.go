package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

type Config struct {
	Port string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required!")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)
	address := "localhost:" + b.config.Port
	log.Println("Starting server on port", address)
	if err := http.ListenAndServe(address, handler); err != nil {
		log.Println("Error starting server:", err)
	} else {
		log.Fatalf("Server stopped!")
	}
}
