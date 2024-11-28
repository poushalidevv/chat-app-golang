package models

import "github.com/jinzhu/gorm"

type Message struct {
    gorm.Model
    ConversationID uint   `json:"conversation_id"`
    SenderID       uint   `json:"sender_id"`       
    RecipientID    uint   `json:"recipient_id"`    
    Content        string `json:"content"`
    Timestamp      int64  `json:"timestamp"` 
}
