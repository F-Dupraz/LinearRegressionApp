package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/F-Dupraz/LinearRegressionApp/server"
)

type ResponseRoot struct {
	Message string `json:"message"`
	Status bool `json:"status"`
}

func HandlerRoot(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseRoot {
			Message: "So far so good!",
			Status: true,
		})
	}
}
