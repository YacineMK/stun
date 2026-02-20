package internal

import (
	"log"
	"net"
)

func HandlePacket(buf []byte, conn net.PacketConn) {
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println("Error reading from connection:", err)
			continue
		}
		log.Printf("Received %d bytes from %s: %s\n", n, addr.String(), string(buf[:n]))
	}
}
