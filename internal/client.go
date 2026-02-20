package internal

import (
	"log"
	"net"

	"github.com/YacineMK/stun/utils"
)

func ClientHandler(serverAddr string, conn net.Conn) {

	req := utils.BuildBindingRequest()
	_, err := conn.Write(req[:])
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1500)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	if n < 32 {
		log.Fatal("Response too short")
	}

	ip, port := utils.ReverseXorMappedAddress(buf[24:32])
	log.Fatalf("Public IP: %s, port: %d\n", ip.String(), port)
}
