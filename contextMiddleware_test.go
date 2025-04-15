package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextMiddleware(t *testing.T) {
	// Define test cases
	testCases := []struct {
		queryString string
		expectedID  int
		expectError bool
	}{
		{"id=7", 7, false},      // Valid ID
		{"id=42", 42, false},    // Valid ID
		{"id=invalid", 0, true}, // Invalid ID
		{"", 0, false},          // No ID in query string should not result in an error
	}

	// Iterate over test cases
	for _, tc := range testCases {
		// Create a mock request with the query string
		req := httptest.NewRequest("GET", "/test?"+tc.queryString, nil)
		rr := httptest.NewRecorder()

		// Create a mock handler to verify the context
		mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Retrieve the bunnyID from the context
			bunnyID, ok := r.Context().Value(contextKey("bunnyID")).(int)

			if tc.expectError {
				// If an error is expected, ensure the bunnyID is not set
				if ok {
					t.Errorf("Expected no bunnyID in context, but got %d", bunnyID)
				}
			} else {
				// If no error is expected, ensure the bunnyID matches the expected value
				if !ok && tc.queryString != "" {
					t.Errorf("Expected bunnyID=%d, but it was not set in the context", tc.expectedID)
				} else if ok && bunnyID != tc.expectedID {
					t.Errorf("Expected bunnyID=%d, but got %d", tc.expectedID, bunnyID)
				}
			}
		})

		// Wrap the mock handler with the middleware
		contextMiddleware(mockHandler).ServeHTTP(rr, req)

		// Check for HTTP errors
		if tc.expectError && rr.Code != http.StatusBadRequest {
			t.Errorf("Expected HTTP status %d, but got %d", http.StatusBadRequest, rr.Code)
		} else if !tc.expectError && rr.Code != http.StatusOK {
			t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, rr.Code)
		}
	}
}
