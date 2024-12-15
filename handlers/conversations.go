package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "chat-app-golang/db"
    "chat-app-golang/models"
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/mux"
)

func ListConversationsHandler(w http.ResponseWriter, r *http.Request) {
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

    var conversations []models.Conversation
    if err := db.DB.Where("participants @> ?", []uint{userID}).Find(&conversations).Error; err != nil {
        http.Error(w, "Failed to retrieve conversations", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(conversations)
}

func CreateConversationHandler(w http.ResponseWriter, r *http.Request) {
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

    var conversation models.Conversation
    err = json.NewDecoder(r.Body).Decode(&conversation)
    if err != nil {
        http.Error(w, "Invalid request format", http.StatusBadRequest)
        return
    }

    conversation.Participants = append(conversation.Participants, userID)

    if err := db.DB.Create(&conversation).Error; err != nil {
        http.Error(w, "Failed to create conversation", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(conversation)
}

func GetConversationHandler(w http.ResponseWriter, r *http.Request) {
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

    vars := mux.Vars(r)
    conversationID, err := strconv.ParseUint(vars["conversation_id"], 10, 32)
    if err != nil {
        http.Error(w, "Invalid conversation ID", http.StatusBadRequest)
        return
    }

    var conversation models.Conversation
    if err := db.DB.First(&conversation, uint(conversationID)).Error; err != nil {
        http.Error(w, "Conversation not found", http.StatusNotFound)
        return
    }

    isParticipant := false
    for _, participant := range conversation.Participants {
        if participant == userID {
            isParticipant = true
            break
        }
    }

    if !isParticipant {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var messages []models.Message
    if err := db.DB.Where("conversation_id = ?", conversation.ID).Find(&messages).Error; err != nil {
        http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "conversation": conversation,
        "messages":     messages,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

