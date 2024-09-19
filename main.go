package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
    "sync"
    "time"
)

const defaultPort = "8989"

var (
    clients     = make(map[net.Conn]string)
    clientsLock sync.Mutex
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
        /**
            Read client name
        */
    fmt.Fprintf(conn, "Welcome to TCP-Chat!\n")
    fmt.Fprintf(conn, "Enter your name: ")
        /**
            Read client name
        */
    scanner := bufio.NewScanner(conn)
    if scanner.Scan() {
        name := scanner.Text()
        
        clientsLock.Lock()
        clients[conn] = name
        clientsLock.Unlock()
        /**
            Notify other clients
         */
    
        broadcast(fmt.Sprintf("%s has joined the chat...\n", name))
        /**
            Handle incoming messages
        */ 
        for scanner.Scan() {
            msg := scanner.Text()
            if msg != "" {
                broadcast(fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format("2006-01-02 15:04:05"), name, msg))
            }
        }
            /**
                Notify others when a client leaves
            */
        
        clientsLock.Lock()
        delete(clients, conn)
        clientsLock.Unlock()
        broadcast(fmt.Sprintf("%s has left the chat...\n", name))
    }
}

func broadcast(message string) {
    clientsLock.Lock()
    defer clientsLock.Unlock()
    
    for client := range clients {
        _, err := fmt.Fprint(client, message)
        if err != nil {
            log.Printf("Error broadcasting message: %v", err)
            client.Close()
            delete(clients, client)
        }
    }
}
func main() {
    port := defaultPort
    if len(os.Args) > 1 {
        port = os.Args[1]
    }
    
    listener, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
    defer listener.Close()
    
    log.Printf("Listening on port %s", port)
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %v", err)
            continue
        }
        
        go handleConnection(conn)
    }
}
