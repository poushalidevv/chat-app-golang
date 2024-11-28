package handlers

import (
    "encoding/json"
    "net/http"
    "chat-app-golang/db"
    "chat-app-golang/models"
    "github.com/dgrijalva/jwt-go"
    "log"
)

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
        return
    }

    tokenString := authHeader[len("Bearer "):]

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, http.ErrAbortHandler
        }
        return []byte("your-secret-key"), nil
    })

    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || claims["sub"] == nil {
        http.Error(w, "Invalid token claims", http.StatusUnauthorized)
        return
    }

    subClaim, ok := claims["sub"].(float64)
    if !ok {
        http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
        return
    }
    userID := uint(subClaim)

    log.Printf("User ID from JWT token: %d", userID) 

    var user models.User
    if err := db.DB.First(&user, userID).Error; err != nil {
        log.Printf("Error retrieving user from DB: %v", err) 
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
        return
    }

    tokenString := authHeader[len("Bearer "):]

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, http.ErrAbortHandler
        }
        return []byte("your-secret-key"), nil
    })

    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || claims["sub"] == nil {
        http.Error(w, "Invalid token claims", http.StatusUnauthorized)
        return
    }

    subClaim, ok := claims["sub"].(float64)
    if !ok {
        http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
        return
    }
    userID := uint(subClaim)

    var updatedProfile models.User
    err = json.NewDecoder(r.Body).Decode(&updatedProfile)
    if err != nil {
        http.Error(w, "Invalid request format", http.StatusBadRequest)
        return
    }

    var user models.User
    if err := db.DB.First(&user, userID).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    user.Name = updatedProfile.Name
    user.Username = updatedProfile.Username
    user.Email = updatedProfile.Email
    user.Password = updatedProfile.Password

    if err := db.DB.Save(&user).Error; err != nil {
        http.Error(w, "Failed to update profile", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}
