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
	var name string
	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		name, err = reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		name = strings.TrimSpace(name)
		// Check if the name is empty
		if name == "" {
			conn.Write([]byte("Name is empty. Please choose a name.\n"))
			continue
		}
		// Check if the name is already taken
		mutex.Lock()
		isTaken := false
		for client := range clients {
			if client.name == name {
				conn.Write([]byte("Name already taken. Please choose another one.\n"))
				isTaken = true
				break
			}
		}
		mutex.Unlock()
		if !isTaken {
			break
		}
	}
	mutex.Lock()
	if len(clients) >= maxClients {
		mutex.Unlock()
		conn.Write([]byte("Maximum number of clients reached. Try again later.\n"))
		conn.Close()
		return
	}
	mutex.Unlock()
	client := Client{conn, name}
	clients[client] = true
	message1 := fmt.Sprintf("\n%s has joined our chat...", name)
	for c := range clients {
		if c != client {
			c.conn.Write([]byte(message1))
			c.conn.Write([]byte(fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), c.name)))
		}
	}
	/**
	<!--				comment 				--!>
	*/
	for _, msg := range messages {
		conn.Write([]byte(msg))
	}
	/**
	<!--				comment 				--!>
	*/
	conn.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), name)))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg != "" {
			conn.Write([]byte("\033[1A\033[2K")) // Effacer la ligne dans le terminal du client
			conn.Write([]byte(fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, msg)))
			broadcastMessage(fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, msg), client)
			// Add the message to the list of messages
			mutex.Lock()
			messages = append(messages, fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, msg))
			mutex.Unlock()
		} else {
			conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), name)))
		}
	}
	/**
	<!--		Remove the client from the list of connected clients		--!>
	*/
	delete(clients, client)
	message := fmt.Sprintf("\n%s has left our chat...", name)
	for c := range clients {
		if c != client {
			c.conn.Write([]byte(message))
			c.conn.Write([]byte(fmt.Sprintf("\n[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), c.name)))
		}
	}
	saveMessagesToFile()
	/**
	<!--				 Close the connection with the client				--!>
	*/
	conn.Close()
}
func saveMessagesToFile() error {
	file, err := os.Create("logs.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	mutex.Lock()
	defer mutex.Unlock()
	for _, msg := range messages {
		_, err := file.WriteString(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
