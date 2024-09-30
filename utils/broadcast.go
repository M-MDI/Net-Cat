package utils

import (
	"fmt"
	"time"
)
func broadcastMessage(message string, sender Client) {
	// Blocks access to the shared resource "map clients" to prevent multiple goroutines from simultaneously accessing this resource
	mutex.Lock()
	// Automatically unlocks the resource at the end of the function.
	defer mutex.Unlock()
	/**
	<!--				comment 				--!>
	*/
	// Browse the "clients" map and send the messages to other clients other than the sender by the method (Write) on the TCP connection (client.conn)
	for client := range clients {
		if client != sender {
			client.conn.Write([]byte("\n" + message))
		}
		client.conn.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.name)))
	}
}
