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

func listBunnies() []*Bunny {
	return bunnies
}
