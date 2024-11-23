package repositories

import (
	"chat-app/internal/models"
	"database/sql"
)

type MessageRepository struct {
	DB *sql.DB
}

func (r *MessageRepository) SaveMessage(msg models.Message) error {
	_, err := r.DB.Exec("INSERT INTO messages (content, sender, time) VALUES ($1, $2, $3)",
		msg.Content, msg.Sender, msg.Time)
	return err
}

func (r *MessageRepository) GetMessages() ([]models.Message, error) {
	rows, err := r.DB.Query("SELECT id, content, sender, time FROM messages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.Sender, &msg.Time)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
