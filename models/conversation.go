package models

import "github.com/jinzhu/gorm"

// Conversation model definition
type Conversation struct {
    gorm.Model
    Name         string `json:"name"`       
    Participants []uint `json:"participants"` 
}