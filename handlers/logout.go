package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// In-memory blacklist (for demonstration purposes)
var tokenBlacklist = make(map[string]time.Time)

// LogoutHandler invalidates the JWT token on logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	// Format: Bearer <token>
	tokenString := authHeader[len("Bearer "):]

	// Parse the token to extract claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		// Secret key to validate the token
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Add the token to the blacklist (in-memory map with expiration time)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	// Store the token in the blacklist with an expiration date
	tokenBlacklist[tokenString] = time.Now().Add(24 * time.Hour) // Set expiry to 24 hours

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
