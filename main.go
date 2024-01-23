package main 

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/F-Dupraz/LinearRegressionApp/server"
	"github.com/F-Dupraz/LinearRegressionApp/handlers"
)

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HandlerRoot(s)).Methods(http.MethodGet)

	r.HandleFunc("/predict/inurance", handlers.PredictInsurance(s)).Methods(http.MethodPost)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v.\n", err)
	}

	PORT := os.Getenv("PORT")

	s, err := server.NewServer(context.Background(), &server.Config {
		Port: PORT,
	})

	if err != nil {
		log.Fatalf("Error creating server %v.\n", err)
	}

	s.Start(BindRoutes)
}
