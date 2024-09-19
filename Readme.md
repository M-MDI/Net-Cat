# NetCat Project Plan

## 1. Setup & Initialization
- [ ] Initialize a Go module (`go mod init netcat`).
- [ ] Set up the main file (`main.go`) for your server and client logic.
- [ ] Add basic TCP server functionality that listens on a default port (8989) if no port is specified.

## 2. Implement TCP Server
- [ ] Create a TCP server that listens on the specified port for incoming client connections.
- [ ] Accept multiple client connections using Go routines (1 to many relationship).
- [ ] Handle connections with a channel or Mutex to manage concurrent access.
- [ ] Implement a function to control the number of connections (max 10 clients).

## 3. Client Management
- [ ] Create a function to ask for the client's name when they connect.
- [ ] Ensure that clients with empty names are rejected.
- [ ] Track all connected clients with their names.
- [ ] Notify all clients when a new client joins or leaves the chat.

## 4. Message Handling
- [ ] Implement a function to receive messages from clients.
- [ ] Broadcast messages to all clients, including the timestamp and client name.
- [ ] Ensure that empty messages are not broadcasted.
- [ ] Ensure that clients who leave do not cause others to disconnect.

## 5. Message History
- [ ] Save all messages in a message log.
- [ ] Upload the chat history to any newly joined clients so they can see past messages.

## 6. Error Handling & Usage Message
- [ ] Handle errors when the server or clients fail to connect.
- [ ] Display a proper usage message when the wrong number of arguments is provided (`[USAGE]: ./TCPChat $port`).

## 7. Client-Side Code
- [ ] Implement a client that connects to the TCP server using a given IP and port.
- [ ] Display the Linux logo and request the client’s name upon connection.
- [ ] Ensure the client can send messages to the server.
- [ ] Display messages from other clients with proper formatting (`[timestamp][client.name]: [client.message]`).

## 8. Testing
- [ ] Write unit tests for the server connection and client functionality.
- [ ] Test with multiple clients to verify the group chat functionality.

## 9. Bonus Features (Optional)
- [ ] Implement a terminal UI using the `gocui` package.
- [ ] Save chat logs into a file for future reference.
- [ ] Support multiple group chats running on different ports.

## 10. Final Touches
- [ ] Refactor code to improve readability and maintainability.
- [ ] Review Go best practices (error handling, concurrency, etc.).
- [ ] Create a README explaining how to run the server and client.
- [ ] Perform final testing and debugging before submission.
