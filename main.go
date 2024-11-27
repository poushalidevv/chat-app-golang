package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
	"chat-app-golang/db"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow connections from all origins for simplicity
    },
}

func main() {

	db.InitPostgres()
	db.InitRedis()
    // Set up routes
    http.HandleFunc("/register", RegisterHandler)
    http.HandleFunc("/login", LoginHandler)
    http.HandleFunc("/messages", SendMessageHandler)
    http.HandleFunc("/chat", HandleWebSocket) // WebSocket endpoint

    // Start the server
    fmt.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "User registration handler")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "User login handler")
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Send message handler")
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    fmt.Println("New WebSocket connection established")

    // Handle WebSocket communication here (e.g., read messages, send messages)
    for {
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        // Echo the received message back to the client
        err = conn.WriteMessage(msgType, msg)
        if err != nil {
            log.Println(err)
            return
        }
    }
}
