package main

import (
	"context"
	"net/http"
	"strconv"
)

type contextKey string

// HTTP middleware setting a value on the request context
func contextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			// Convert the ID to an integer
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid Bunny ID", http.StatusBadRequest)
				return
			}
			//Add the bunny ID to the context
			ctx := context.WithValue(r.Context(), contextKey("bunnyID"), id)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
