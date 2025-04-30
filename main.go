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

// Global variable for the database connection
var db *pgx.Conn

// resuseable function to Initialise and connect to the database
func connectToDatabase() (*pgx.Conn, error) {
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
	return conn, nil
}

func setupDatabaseSchema() error {
	// SQL to create the bunnies table if it doesn't exist
	schema := `
    CREATE TABLE IF NOT EXISTS bunnies (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        breed VARCHAR(100) NOT NULL,
        age INT NOT NULL,
        description TEXT,
        characteristics TEXT
    );
    `

	// Execute the schema creation query
	_, err := db.Exec(context.Background(), schema)
	if err != nil {
		return fmt.Errorf("failed to set up database schema: %v", err)
	}

	fmt.Println("Database schema set up successfully!")
	return nil
}

// Insert test data into the bunnies table
func seedDatabase() error {
	seedData := `
    INSERT INTO bunnies (name, breed, age, description, characteristics)
    VALUES
    ('Fluffy', 'Angora', 2, 'A fluffy bunny', 'Friendly and playful'),
    ('Snowball', 'Lionhead', 1, 'A small white bunny', 'Energetic and curious'),
    ('Thumper', 'Dutch', 3, 'A calm bunny', 'Loves to hop around');
    `

	_, err := db.Exec(context.Background(), seedData)
	if err != nil {
		return fmt.Errorf("failed to seed database: %v", err)
	}

	fmt.Println("Database seeded with test data!")
	return nil
}

// main function connects to the databses, initializes the router and starts the server on port 8000.
// It also sets up a simple route that responds with a welcome message.
func main() {
	// Connect to the database
	var err error
	db, err = connectToDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// Set up the database schema
	err = setupDatabaseSchema() //function designed to return an error value. Tells us if it was successful or not.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up database schema: %v\n", err)
		os.Exit(1)
	}

	// Seed the database with test data
	err = seedDatabase() //calls the seedDatabase function to insert test data into the bunnies table
	if err != nil {      //check if the seedDatabase function returns an error
		fmt.Fprintf(os.Stderr, "Error seeding database: %v\n", err) //if there is an error, print it to stderr
		// and exit the program with a non-zero status code
		os.Exit(1) // exit the program with a non-zero status code
	}

	// Test the connection by executing a simple query to return welcome message
	var greeting string
	err = db.QueryRow(context.Background(), "select 'Welcome to GoBunnyAPI!'").Scan(&greeting)
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
	// r.Get("/test/bunnies", bunnyHandler.ListBunnies) // Route to get all bunnies
	// r.Get("/test/bunny", bunnyHandler.GetBunnyByID)  // Route to get a specific bunny by ID
	r.Get("/test/bunnies", bunnyHandler.GetBunnies)     // Route to get all bunnies from the database
	r.Post("/test/makeBunny", bunnyHandler.CreateBunny) // Route to create a new bunny
	return r
}
