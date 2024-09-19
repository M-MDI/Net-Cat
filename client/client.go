package NetCat
import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run client.go <server-address>:<port>")
        return
    }
    
    serverAddress := os.Args[1]
    conn, err := net.Dial("tcp", serverAddress)
    if err != nil {
        fmt.Printf("Error connecting to server: %v\n", err)
        return
    }
    defer conn.Close()
    
    // Handle incoming messages
    go func() {
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
    }()
    
    // Send messages to server
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        _, err := fmt.Fprintln(conn, scanner.Text())
        if err != nil {
            fmt.Printf("Error sending message: %v\n", err)
            return
        }
    }
}
