package main

import (
	"log"
	"net"

	"github.com/YacineMK/stun/internal"
)

func main() {
	conn, _ := net.ListenPacket("udp", ":3478")
	defer conn.Close()

	log.Println("Starting STUN Server")
	buf := make([]byte, 1500)
	go internal.PacketHandler(buf, conn)

	select {}
}
