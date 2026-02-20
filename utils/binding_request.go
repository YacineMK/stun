package utils

import (
	"encoding/binary"
	"math/rand"
	"time"
)

func BuildBindingRequest() [20]byte {
	var req [20]byte

	req[0] = 0x00
	req[1] = 0x01
	req[2], req[3] = 0x00, 0x00
	binary.BigEndian.PutUint32(req[4:8], 0x2112A442)

	rand.Seed(time.Now().UnixNano())
	for i := 8; i < 20; i++ {
		req[i] = byte(rand.Intn(256))
	}

	return req
}
