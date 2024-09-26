# TCP Chat

## Project Overview
TCP Chat is a server-client architecture project that recreates the functionality of NetCat (nc) in Go. It implements a group chat system where multiple clients can connect to a server and exchange messages in real-time.

## Features
- TCP connection between server and multiple clients (1-to-many relationship)
- Client name requirement for identification
- Connection quantity control (maximum 10 connections)
- Message broadcasting to all connected clients
- Time-stamped messages with sender identification
- New clients receive chat history upon joining
- Join and leave notifications for all clients
- Automatic port selection (default 8989) with custom port option

## Technologies Used
- Go (Golang)
- TCP/IP Networking
- Goroutines for concurrent operations
- Channels or Mutexes for synchronization

## Usage
To start the server:
```
go run . [port]
```
If no port is specified, the server will listen on port 8989.

To connect as a client:
```
nc <server_ip> <port>
```

## Project Structure
- Server implementation in Go
- Client handling and message broadcasting
- Error handling for both server and client sides

## Development
This project was developed as part of a programming exercise to understand and implement network programming concepts in Go.

Developed by: MorningStar
