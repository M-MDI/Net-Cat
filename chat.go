package main
import (
	"fmt"
	"log"
	"net"
	"os"
	"netcat/utils"
)
func main() {
	port := ""
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 1 {
		port = "8989"
	} else {
		port = os.Args[1]
	}
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Print("Listening on the port :", port, "\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go utils.HandleConnection(conn)
	}
}
