# Chat App Backend
This is the backend for a chat application built with Go. It provides user authentication, profile management, and real-time chat functionality using WebSockets and Redis.
## Table of Contents
- Installation
- Configuration
- Usage
- Endpoints
  - [User Registration](#user-registration)
  - [User Login](#user-login)
  - [User Logout](#user-logout)
  - [Profile Management](#profile-management)
  - [WebSocket Chat](#websocket-chat)
- License
## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/chat-app-golang.git
    cd chat-app-golang
    ```
2. Install dependencies:
    ```sh
    go mod download
    ```
3. Set up the environment variables by creating a .env file in the root directory:
    ```env
    SERVER_ADDRESS=:8080
    REDIS_ADDRESS=localhost:6379
    DATABASE_URL=postgres://postgres:Aviral123@localhost:5432/chat_app
    ```
4. Run the database migrations:
    ```sh
    go run main.go
    ```
## Configuration
Ensure that the .env file is correctly set up with the necessary environment variables for PostgreSQL and Redis connections.
## Usage
1. Build and run the application:
    ```sh
    go build -o main .
    ./main
    ```
2. The server will start on http://localhost:8080.
## Endpoints
### User Registration
- URL: /register
- Method: POST
- Request Body:
    ```json
    {
        "name": "John Doe",
        "username": "johndoe",
        "email": "johndoe@example.com",
        "password": "password123"
    }
    ```
- Response:
    ```json
    {
        "id": 1,
        "name": "John Doe",
        "username": "johndoe",
        "email": "johndoe@example.com",
        "createdAt": "2023-10-01T12:00:00Z",
        "updatedAt": "2023-10-01T12:00:00Z"
    }
    ```
### User Login
- URL: /login
- Method: POST
- Request Body:
    ```json
    {
        "usernameOrEmail": "johndoe",
        "password": "password123"
    }
    ```
- Response:
    ```json
    {
        "message": "Login successful",
        "token": "jwt-token"
    }
    ```
### User Logout
- URL: /logout
- Method: POST
- Headers:
    - Authorization: Bearer <jwt-token>
- Response:
    ```json
    {
        "message": "Logout successful"
    }
    ```
### Profile Management
#### Get Profile
- URL: /profile
- Method: GET
- Headers:
    - Authorization: Bearer <jwt-token>
- Response:
    ```json
    {
        "id": 1,
        "name": "John Doe",
        "username": "johndoe",
        "email": "johndoe@example.com",
        "createdAt": "2023-10-01T12:00:00Z",
        "updatedAt": "2023-10-01T12:00:00Z"
    }
    ```
#### Update Profile
- URL: /profile
- Method: PUT
- Headers:
    - Authorization: Bearer <jwt-token>
- Request Body:
    ```json
    {
        "name": "John Doe Updated",
        "username": "johndoeupdated",
        "email": "johndoeupdated@example.com",
        "password": "newpassword123"
    }
    ```
- Response:
    ```json
    {
        "id": 1,
        "name": "John Doe Updated",
        "username": "johndoeupdated",
        "email": "johndoeupdated@example.com",
        "createdAt": "2023-10-01T12:00:00Z",
        "updatedAt": "2023-10-01T12:30:00Z"
    }
    ```

### WebSocket Chat
- URL: /chat
- Method: GET
- Headers:
    - Authorization: Bearer <jwt-token>
- Description: Establishes a WebSocket connection for real-time chat.
#### WebSocket Communication
Once the WebSocket connection is established, the client can send and receive messages in real-time. The server listens for messages from the client, publishes them to a Redis channel, and broadcasts messages from the Redis channel to all connected WebSocket clients.
**Example WebSocket Message:**
```json
{
    "sender_id": 1,
    "recipient_id": 2,
    "content": "Hello, how are you?",
    "timestamp": 1633024800
}
```
**Handling WebSocket Messages:**
1. **Client to Server:**
   - The client sends a message to the server via the WebSocket connection.
   - The server reads the message and publishes it to the Redis channel.
2. **Server to Client:**
   - The server listens for messages on the Redis channel.
   - When a message is received, the server broadcasts it to all connected WebSocket clients.
## Contact
For any questions or inquiries, please contact me at [yaviral08@gmail.com](mailto:yaviral08@@gmail.com).

