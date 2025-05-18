package rest

import (
	"log/slog"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	slog.Info("Hello World")
}

func NewRestAPI() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/func", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		slog.Info("Func Call Path")
	})

	mux.Handle("/api", apiHandler{})

	http.ListenAndServe(":8080", mux)
}
