package main

import "net"

func main() {
	conn, err := net.Dial("udp", "localhost:3478")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			panic(err)
		}
		println(string(bytes[:n]))
	}
}
