package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "github.com/gorilla/mux"
    "chat-app-golang/db"
    "chat-app-golang/handlers"
    "github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow connections from all origins for simplicity
    },
}

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    db.InitPostgres()
    db.InitRedis()

    // Create a new router
    r := mux.NewRouter()

    // Set up routes
    r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
    r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

    // Profile handling with GET and PUT methods
    r.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            handlers.GetProfileHandler(w, r)  // GET: Retrieve user profile
        case "PUT":
            handlers.UpdateProfileHandler(w, r) // PUT: Update user profile
        default:
            http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
        }
    })

    // Messaging routes with method-specific handlers
    r.HandleFunc("/messages", handlers.SendMessageHandler).Methods("POST")
    r.HandleFunc("/messages", handlers.GetMessagesHandler).Methods("GET")

    r.HandleFunc("/conversations", handlers.ListConversationsHandler).Methods("GET")
    r.HandleFunc("/conversations", handlers.CreateConversationHandler).Methods("POST")
    r.HandleFunc("/conversations/{conversation_id}", handlers.GetConversationHandler).Methods("GET")
    // WebSocket route
    r.HandleFunc("/chat", handlers.HandleWebSocket)

    // Start the server
    fmt.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

// WebSocket handler
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
        fmt.Println("Waiting for message...")
        msgType, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            log.Println(err)
            return
        }
        fmt.Printf("Received message: %s\n", msg)
        // Echo the received message back to the client
        err = conn.WriteMessage(msgType, msg)
        if err != nil {
            fmt.Println("Error writing message:", err)
            log.Println(err)
            return
        }
    }
}
