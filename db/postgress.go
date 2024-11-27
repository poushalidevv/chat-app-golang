package db

import (
    "context"
    "fmt"
    "log"
    "github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func InitPostgres() {
    var err error
    Conn, err = pgx.Connect(context.Background(), "postgres://postgres:Aviral123@localhost:5432/chat_app")
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    fmt.Println("Successfully connected to PostgreSQL!")
}
