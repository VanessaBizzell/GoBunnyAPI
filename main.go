package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// main function initializes the router and starts the server on port 8000.
// It also sets up a simple route that responds with a welcome message.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(contextMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to GoBunnyAPI!"))
	})

	r.Mount("/api/v1", bunnyRoutes()) // Mount the bunny routes on the root path

	http.ListenAndServe(":8000", r)
}

// return `localhost:8000/api/v1/test/bunnies`
func bunnyRoutes() chi.Router {
	r := chi.NewRouter()
	bunnyHandler := BunnyHandler{}
	r.Get("/test/bunnies", bunnyHandler.ListBunnies) // Route to get all bunnies
	r.Get("/test/bunny", bunnyHandler.GetBunnyByID)  // Route to get a specific bunny by ID
	return r
}
