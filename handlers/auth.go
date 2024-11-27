package handlers

import (
	"encoding/json"
	"net/http"
	"chat-app-golang/db"
	"chat-app-golang/models"
	"chat-app-golang/utils"
)

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse the incoming JSON request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before saving
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	// Save the new user to the database
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	// Return the created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
