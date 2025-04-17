package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// main function connects to the databses, initializes the router and starts the server on port 8000.
// It also sets up a simple route that responds with a welcome message.
func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	//connect to the database using pgx
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Welcome to GoBunnyAPI!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	// Initialize the router
	// Use chi for routing and middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(romanToBunnyID)
	r.Use(contextMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GoBunnyAPI! router is working!"))
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
