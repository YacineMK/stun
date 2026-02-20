package internal

import (
	"log"
	"net"

	"github.com/YacineMK/stun/utils"
)

func PacketHandler(buf []byte, conn net.PacketConn) {
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println("read error:", err)
			continue
		}

		if n > 1500 {
			log.Println("message too long from", addr.String())
			conn.WriteTo([]byte("Error: Message too long"), addr)
			continue
		}

		var header [20]byte
		copy(header[:], buf[:20])
		req := NewMessage(header)

		if req.Type != 0x0001 {
			log.Println("unsupported message type from", addr.String())
			conn.WriteTo([]byte("Error: Unsupported message type"), addr)
			continue
		}

		if req.MagicCookie != 0x2112A442 {
			log.Println("invalid magic cookie from", addr.String())
			conn.WriteTo([]byte("Error: Invalid magic cookie"), addr)
			continue
		}

		log.Println("valid binding request from", addr.String())

		resp := &Message{
			Type:        0x0101,
			Length:      0,
			MagicCookie: 0x2112A442,
			TxID:        req.TxID,
		}

		respBytes := resp.Marshal()
		attr := utils.XorMappedAddress(addr)
		resp.Length = uint16(len(attr))
		respBytes = resp.Marshal()

		packet := append(respBytes[:], attr...)
		_, err = conn.WriteTo(packet, addr)
		if err != nil {
			log.Println("send error:", err)
		}
	}
}
