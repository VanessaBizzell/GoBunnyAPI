package main

import (
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

// function to list all bunnies
func listBunnies() []*Bunny {
	return bunnies
}
