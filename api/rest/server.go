package rest

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	slog.Info("Hello World")
}

type Payload struct {
	Message string
}

func (apiHandler) GetAPIFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		body := make(map[string]string)
		body["message"] = "Hello World"

		payload, e := json.Marshal(body)
		if e != nil {
			slog.Error("Issue marshalling payload")
			w.WriteHeader(http.StatusInternalServerError)
		}

		_, e = w.Write(payload)
		if e != nil {
			slog.Error("Issue writing body to response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			slog.Error("Unable to read user request")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var payload Payload
			err := json.Unmarshal(body, &payload)
			if err != nil {
				slog.Error("Issue unmarshaling payload")
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				slog.Info("User message:", "message", payload.Message)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func NewRestAPI() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/func", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		slog.Info("Func Call Path")
	})

	apiHandler := &apiHandler{}
	mux.Handle("/api", apiHandler)
	mux.HandleFunc("/api/handler", apiHandler.GetAPIFunc)

	http.ListenAndServe(":8080", mux)
}
