package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(1 * time.Second))
	r.Get("/hello", getHandler)
	r.Get("/bad", badHandler)
	r.Get("/panic", panicHandler)
	r.Get("/timeout", timeoutHandler)
	r.Post("/echo", betterEndpoint)

	http.ListenAndServe(":8080", r)
}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
}

func panicHandler(_ http.ResponseWriter, _ *http.Request) {
	panic("oh no!")
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func betterEndpoint(w http.ResponseWriter, r *http.Request) {
	_, err := io.Copy(w, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}
