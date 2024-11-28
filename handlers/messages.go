package handlers

import (
	"encoding/json"
	"net/http"
	"chat-app-golang/db"
	"chat-app-golang/models"
	"log"
	"strconv"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var newMessage models.Message

	log.Printf("Received request to send a message: %v", r)

	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		log.Printf("Error decoding message request: %v", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	log.Printf("Parsed message: %+v", newMessage)

	if newMessage.ConversationID == 0 || newMessage.SenderID == 0 || newMessage.Content == "" {
		log.Printf("Missing required fields in the message request: %+v", newMessage)
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&newMessage).Error; err != nil {
		log.Printf("Error saving message to database: %v", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	log.Printf("Message sent successfully: %+v", newMessage)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMessage)
}

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request to fetch messages for conversation ID: %v", r.URL.Query().Get("conversation_id"))

	conversationIDStr := r.URL.Query().Get("conversation_id")
	if conversationIDStr == "" {
		log.Println("Missing conversation ID in the request")
		http.Error(w, "Missing conversation ID", http.StatusBadRequest)
		return
	}

	conversationID, err := strconv.Atoi(conversationIDStr)
	if err != nil || conversationID <= 0 {
		log.Printf("Invalid conversation ID: %v", conversationIDStr)
		http.Error(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	var messages []models.Message
	if err := db.DB.Where("conversation_id = ?", conversationID).Find(&messages).Error; err != nil {
		log.Printf("Error retrieving messages for conversation ID %d: %v", conversationID, err)
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}

	log.Printf("Retrieved %d messages for conversation ID %d", len(messages), conversationID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}
