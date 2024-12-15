package handlers

import (
    "log"
    "chat-app-golang/db"
    "context"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        // Implement your origin check logic (allow all origins for simplicity)
        return true
    },
}

// HandleWebSocket handles WebSocket connections for real-time chat
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    // Upgrade HTTP to WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket Upgrade Error:", err)
        return
    }
    defer conn.Close()

    // Subscribe to the "chat" channel
    pubsub := db.RedisClient.Subscribe(context.Background(), "chat")
    defer pubsub.Close()

    // Channel to receive messages from Redis
    redisMessages := pubsub.Channel()

    // Channel to signal when the connection is closed
    done := make(chan struct{})

    // Goroutine to read from WebSocket client and publish to Redis
    go func() {
        defer close(done)
        for {
            _, msg, err := conn.ReadMessage()
            if err != nil {
                log.Println("WebSocket Read Error:", err)
                return
            }
            // Publish the message to Redis
            err = db.RedisClient.Publish(context.Background(), "chat", msg).Err()
            if err != nil {
                log.Println("Redis Publish Error:", err)
                return
            }
        }
    }()

    // Read messages from Redis and send to WebSocket client
    for {
        select {
        case msg := <-redisMessages:
            // Send message to WebSocket client
            err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
            if err != nil {
                log.Println("WebSocket Write Error:", err)
                return
            }
        case <-done:
            // Exit the loop if the connection is closed
            return
        }
    }
}
