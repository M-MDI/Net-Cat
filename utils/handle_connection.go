package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func HandleConnection(conn net.Conn) {
	// Prints a welcome message as well as the logo via the TCP connection between the client and the server (net.Conn)
	conn.Write([]byte("Welcome to TCP-Chat!\n"))

	file, err := os.Open("logo.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanned := bufio.NewScanner(file)
	for scanned.Scan() {
		conn.Write([]byte(scanned.Text() + "\n"))
	}
	// read customer input(NewReader) to end of line(ReadString)
	reader := bufio.NewReader(conn)
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Print(err)
		return
	}

	name = strings.TrimSpace(name) // delete the name after saving it
	if name == "" || clients[Client{nil, name}] {
		conn.Write([]byte("Name already taken or empty. Please choose another one.\n"))
		conn.Close()
		return
	}

	client := Client{conn, name}
	clients[client] = true

	message1 := fmt.Sprintf("%s has joined our chat...\n", name)
	for c := range clients {
		if c != client {
			c.conn.Write([]byte(message1))
		}
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg != "" {
			conn.Write([]byte("\033[1A\033[2K")) // Effacer la ligne dans le terminal du client
			conn.Write([]byte(fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, msg)))
			broadcastMessage(fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, msg), client)
		}
	}

	//Remove the client from the list of connected clients
	delete(clients, client)

	message := fmt.Sprintf("%s has left our chat...\n", name)
	for c := range clients {
		if c != client {
			c.conn.Write([]byte(message))
		}
	}
	// Close the connection with the client
	conn.Close()
}
