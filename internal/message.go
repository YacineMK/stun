package internal

import (
	"encoding/binary"
)

type Message struct {
	Type        uint16
	Length      uint16
	MagicCookie uint32
	TxID        [12]byte
}

func NewMessage(m [20]byte) *Message {
	return &Message{
		Type:        binary.BigEndian.Uint16(m[0:2]),
		Length:      binary.BigEndian.Uint16(m[2:4]),
		MagicCookie: binary.BigEndian.Uint32(m[4:8]),
		TxID:        [12]byte{m[8], m[9], m[10], m[11], m[12], m[13], m[14], m[15], m[16], m[17], m[18], m[19]},
	}
}

func (m *Message) Marshal() [20]byte {
	var buf [20]byte
	binary.BigEndian.PutUint16(buf[0:2], m.Type)
	binary.BigEndian.PutUint16(buf[2:4], m.Length)
	binary.BigEndian.PutUint32(buf[4:8], m.MagicCookie)
	copy(buf[8:], m.TxID[:])
	return buf
}
