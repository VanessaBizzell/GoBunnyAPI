package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type BunnyHandler struct {
}

// encodes bunny data into JSON format and writes it to the response writer.
func (b BunnyHandler) ListBunnies(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(listBunnies())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

// function to get a specific bunny by ID
func (b BunnyHandler) GetBunnyByID(w http.ResponseWriter, r *http.Request) {
	// Retrieve the bunny ID from the context
	bunnyID, ok := r.Context().Value(contextKey("bunnyID")).(int)
	if !ok {
		http.Error(w, "Bunny ID not provided", http.StatusBadRequest)
		return
	}
	// Find the bunny by ID
	for _, bunny := range bunnies {
		if bunny.ID == bunnyID {
			json.NewEncoder(w).Encode(bunny)
			return
		}
	}
	http.Error(w, "Bunny not found", http.StatusNotFound)
}

// // function to list all bunnies
// func listBunnies() []*Bunny {
// 	return bunnies
// }

// function to list all bunnies from the database
func (b BunnyHandler) GetBunnies(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(), "SELECT id, name, breed, age, description, characteristics FROM bunnies")
	if err != nil {
		http.Error(w, "No bunnies found", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// Define a slice to hold the bunnies
	var bunnies []map[string]interface{}
	// Iterate over the rows and process each bunny
	for rows.Next() {
		var id int
		var name, breed, description, characteristics string
		var age int
		// Scan the row into variables
		err := rows.Scan(&id, &name, &breed, &age, &description, &characteristics)
		if err != nil {
			http.Error(w, "Failed to parse bunny data", http.StatusInternalServerError)
			return
		}
		// Append the bunny to the slice
		bunnies = append(bunnies, map[string]any{
			"id":              id,
			"name":            name,
			"breed":           breed,
			"age":             age,
			"description":     description,
			"characteristics": characteristics,
		})
	}
	// Check for errors after iterating over rows
	if err = rows.Err(); err != nil {
		http.Error(w, "Error reading bunnies from database", http.StatusInternalServerError)
		return
	}
	// Check if no entries are found
	if len(bunnies) == 0 {
		http.Error(w, "No bunnies found", http.StatusNotFound)
		return
	}
	// Encode the bunnies slice as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bunnies)
	if err != nil {
		http.Error(w, "Failed to encode bunnies as JSON", http.StatusInternalServerError)
		return
	}
}

// function to create a new bunny
func (b BunnyHandler) CreateBunny(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the new bunny data
	var newBunny Bunny
	err := json.NewDecoder(r.Body).Decode(&newBunny)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Insert the new bunny into the database
	_, err = db.Exec(context.Background(),
		"INSERT INTO bunnies (name, breed, age, description, characteristics) VALUES ($1, $2, $3, $4, $5)",
		newBunny.Name, newBunny.Breed, newBunny.Age, newBunny.Description, newBunny.Characteristics)
	if err != nil {
		http.Error(w, "Failed to create bunny", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBunny)
}
