package main

import (
	"net/http"
	"strconv"
	"strings"
)

// convert roman numerals to integers
func romanToInt(s string) int {
	// Validate the input
	if !isValidRoman(s) {
		return 0
	}
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i, char := range s {
		if i < len(s)-1 && roman[char] < roman[rune(s[i+1])] {
			result -= roman[char]
		} else {
			result += roman[char]
		}
	}
	return result
}

func isValidRoman(s string) bool {
	if s == "" {
		return false
	}
	validRoman := "IVXLCDM"
	for _, char := range s {
		if !strings.ContainsRune(validRoman, char) {
			return false
		}
	}
	return true
}

// if query string includes roman numerals at end of string then convert the string to int
// once converted, return the int value
// treat the int value as a bunny ID and return the bunny with that ID
func romanToBunnyID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the last part of the URL path
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) > 0 {
			romanPart := parts[len(parts)-1] // Get the last segment of the path
			if isValidRoman(romanPart) {
				id := romanToInt(romanPart) // Convert Roman numeral to integer
				// Rewrite the URL to include the query parameter
				r.URL.Path = "/api/v1/test/bunny"
				q := r.URL.Query()
				q.Set("id", strconv.Itoa(id))
				r.URL.RawQuery = q.Encode()
				// log.Printf("romanToBunnyID: Rewrote URL to %s?id=%d\n", r.URL.Path, id)
			}
		}
		next.ServeHTTP(w, r)
	})
}
