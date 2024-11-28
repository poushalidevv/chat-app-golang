package db

import (
	"context"
	"fmt"
	"log"
	"chat-app-golang/models"
	"github.com/jinzhu/gorm"
	"github.com/jackc/pgx/v4"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Conn *pgx.Conn    // pgx connection for raw SQL interactions
var DB *gorm.DB       // GORM DB object for ORM and AutoMigrate


func InitPostgres() {
	var err error

	
	Conn, err = pgx.Connect(context.Background(), "postgres://postgres:Aviral123@localhost:5432/chat_app")
	if err != nil {
		log.Fatalf("Unable to connect to database using pgx: %v", err)
	}
	fmt.Println("Successfully connected to PostgreSQL using pgx!")

	DB, err = gorm.Open("postgres", "postgres://postgres:Aviral123@localhost:5432/chat_app?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to PostgreSQL using GORM: %v", err)
	}
	fmt.Println("Successfully connected to PostgreSQL using GORM!")

	
	DB.AutoMigrate(&models.User{}) 
	DB.AutoMigrate(&models.Message{})
	fmt.Println("Database schema migrated!")
}
