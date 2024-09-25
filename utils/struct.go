package utils

import (
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
}

var (
	clients = make(map[Client]bool)
	mutex   = &sync.Mutex{}
)
