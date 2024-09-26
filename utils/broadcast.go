package utils

func broadcastMessage(message string, sender Client) {
	// Blocks access to the shared resource "map clients" to prevent multiple goroutines from simultaneously accessing this resource
	mutex.Lock()
	//Automatically unlocks the resource at the end of the function.
	defer mutex.Unlock()
	/**
		<!--				comment 				--!>
	*/
	// Browse the "clients" map and send the messages to other clients other than the sender by the method (Write) on the TCP connection (client.conn).
	for client := range clients {
		if client != sender {
			client.conn.Write([]byte(message))
		}
	}
}
