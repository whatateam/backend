package main

import (
	"net/http"
	"os"
	"cmd/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Implemented"))
	})

	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/users", user.UsersHandler).Methods("GET")
	r.Handle("/teams", NotImplemented).Methods("GET")
	r.Handle("/projects", NotImplemented).Methods("GET")
	r.Handle("/tasks", NotImplemented).Methods("GET")

	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})
