package utils

import (
	"encoding/binary"
	"net"
)

func XorMappedAddress(addr net.Addr) []byte {
	udpAddr := addr.(*net.UDPAddr)
	ip := udpAddr.IP.To4()
	if ip == nil {
		return nil
	}

	b := make([]byte, 12)
	b[0] = 0x00
	b[1] = 0x20
	b[2] = 0x00
	b[3] = 0x08

	b[4] = 0x00
	b[5] = 0x01

	port := uint16(udpAddr.Port) ^ 0x2112
	b[6] = byte(port >> 8)
	b[7] = byte(port & 0xFF)

	magicCookie := []byte{0x21, 0x12, 0xA4, 0x42}
	for i := 0; i < 4; i++ {
		b[8+i] = ip[i] ^ magicCookie[i]
	}
	return b
}

func ReverseXorMappedAddress(attr []byte) (net.IP, uint16) {
	if len(attr) < 8 {
		return nil, 0
	}

	xorPort := binary.BigEndian.Uint16(attr[2:4]) ^ 0x2112
	ip := make(net.IP, 4)
	magicCookie := []byte{0x21, 0x12, 0xA4, 0x42}
	for i := 0; i < 4; i++ {
		ip[i] = attr[4+i] ^ magicCookie[i]
	}
	return ip, xorPort
}
