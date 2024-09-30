package utils
import (
	"net"
	"sync"
)
/*
*
<!--				comment 				--!>
*/
type Client struct {
	conn net.Conn
	name string
}
/*
*
<!--				comment 				--!>
*/
var (
	clients    = make(map[Client]bool)
	messages   = []string{}
	mutex      = &sync.Mutex{}
	maxClients = 10
)