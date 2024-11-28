package handlers

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "chat-app-golang/db"
    "chat-app-golang/models"
    "chat-app-golang/utils"
    "github.com/dgrijalva/jwt-go"
    "time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var userLogin models.UserLogin

    err := json.NewDecoder(r.Body).Decode(&userLogin)
    if err != nil {
        http.Error(w, "Invalid request format", http.StatusBadRequest)
        return
    }

    fmt.Println("Attempting login for:", userLogin.UsernameOrEmail)

    var user models.User
    if err := db.DB.Where("username = ? OR email = ?", userLogin.UsernameOrEmail, userLogin.UsernameOrEmail).First(&user).Error; err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    fmt.Println("User found:", user.Username)

    if !utils.CheckPasswordHash(userLogin.Password, user.Password) {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    token, err := generateJWT(user)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "message": "Login successful",
        "token":   token,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) 
    fmt.Println(response)
    json.NewEncoder(w).Encode(response)
}

func generateJWT(user models.User) (string, error) {
    log.Println("User ID from database:", user.ID) 
    claims := jwt.MapClaims{
        "sub":      user.ID,                  
        "exp":      time.Now().Add(24 * time.Hour).Unix(), 
        "iat":      time.Now().Unix(),         
        "username": user.Username,             
        "email":    user.Email,                
    }

    secretKey := []byte("your-secret-key") 

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
