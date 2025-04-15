package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	// Define test cases
	testCases := []struct {
		roman    string
		expected int
	}{
		{"I", 1},
		{"IV", 4},
		{"VII", 7},
		{"XII", 12},
		{"XX", 20},
		{"XL", 40},
		{"XC", 90},
		{"C", 100},
		{"CD", 400},
		{"CM", 900},
		{"MCMXCIV", 1994},
		{"INVALID", 0}, // Invalid input should return 0
	}

	// Iterate over test cases
	for _, tc := range testCases {
		result := romanToInt(tc.roman)
		if result != tc.expected {
			t.Errorf("romanToInt(%q) = %d; want %d", tc.roman, result, tc.expected)
		}
	}
}

func TestIsValidRoman(t *testing.T) {
	// Define test cases
	testCases := []struct {
		roman    string
		expected bool
	}{
		{"I", true},
		{"IV", true},
		{"VII", true},
		{"XII", true},
		{"INVALID", false}, // Contains invalid characters
		{"123", false},     // Contains numbers
		{"", false},        // Empty string
		{"IIII", true},     // Valid Roman numeral (though unconventional)
		{"ABC", false},     // Contains invalid characters
	}

	// Iterate over test cases
	for _, tc := range testCases {
		result := isValidRoman(tc.roman)
		if result != tc.expected {
			t.Errorf("isValidRoman(%q) = %v; want %v", tc.roman, result, tc.expected)
		}
	}
}

func TestRomanToBunnyID(t *testing.T) {
	// Define test cases
	testCases := []struct {
		inputURL       string
		expectedPath   string
		expectedQuery  string
		expectedStatus int
	}{
		{"/api/v1/test/I", "/api/v1/test/bunny", "id=1", http.StatusOK},     // Valid Roman numeral
		{"/api/v1/test/V", "/api/v1/test/bunny", "id=5", http.StatusOK},     // Valid Roman numeral
		{"/api/v1/test/X", "/api/v1/test/bunny", "id=10", http.StatusOK},    // Valid Roman numeral
		{"/api/v1/test/INVALID", "/api/v1/test/INVALID", "", http.StatusOK}, // Invalid Roman numeral
	}

	// Iterate over test cases
	for _, tc := range testCases {
		// Log the starting URL
		t.Logf("Starting URL: %s", tc.inputURL)

		// Create a mock request
		req := httptest.NewRequest("GET", tc.inputURL, nil)
		rr := httptest.NewRecorder()

		// Create a mock handler to verify the rewritten URL
		mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Print the resulting URL
			t.Logf("Resulting URL: %s", r.URL.String())

			// Check the rewritten path
			if r.URL.Path != tc.expectedPath {
				t.Errorf("Expected path %q, got %q", tc.expectedPath, r.URL.Path)
			}
			// Check the rewritten query string
			if r.URL.RawQuery != tc.expectedQuery {
				t.Errorf("Expected query %q, got %q", tc.expectedQuery, r.URL.RawQuery)
			}
			w.WriteHeader(http.StatusOK)
		})

		// Wrap the mock handler with the middleware
		romanToBunnyID(mockHandler).ServeHTTP(rr, req)

		// Check the response status
		if rr.Code != tc.expectedStatus {
			t.Errorf("Expected status %d, got %d", tc.expectedStatus, rr.Code)
		}
	}
}
