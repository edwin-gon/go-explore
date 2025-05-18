package rest

import "net/http"

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	print("Hello World")
}

func NewRestAPI() {
	http.HandleFunc("/api/func", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		print("Func call!")
	})

	http.Handle("/api", apiHandler{})

	http.ListenAndServe(":8080", nil)
}
