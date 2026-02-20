package main

import (
	"log"
	"net"
	"os"

	"github.com/YacineMK/stun/internal"
)

func main() {
	serverAddr := "127.0.0.1:3478" 
	if len(os.Args) > 1 {
		serverAddr = os.Args[1]
	}

	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	internal.ClientHandler(serverAddr, conn)
}