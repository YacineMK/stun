# stun

A minimal STUN server and client implementation in Go, following [RFC 5389](https://datatracker.ietf.org/doc/html/rfc5389).

## What is STUN?

STUN (Session Traversal Utilities for NAT) is a protocol that allows a client to discover its public IP address and port as seen from the outside network. It is commonly used in WebRTC and peer-to-peer applications to establish direct connections through NATs and firewalls.

## How it works

1. The **client** sends a STUN Binding Request to the server over UDP.
2. The **server** validates the request (magic cookie, message type) and replies with a Binding Response containing the client's public IP and port encoded as a `XOR-MAPPED-ADDRESS` attribute.
3. The **client** decodes the response and prints its public IP and port.

## Project Structure

```
stun/
├── cmd/
│   ├── server/main.go   # STUN server entrypoint
│   └── client/main.go   # STUN client entrypoint
├── internal/
│   ├── server.go        # Packet handler (request validation + response)
│   ├── client.go        # Client handler (send request + decode response)
│   └── message.go       # STUN message types and marshalling
└── utils/
    └── xor_mapped.go    # XOR-MAPPED-ADDRESS encode/decode
```

## Requirements

- Go 1.21+

## Running

**Start the server:**
```bash
go run ./cmd/server/main.go
```

The server listens on UDP port `3478` by default.

**Run the client:**
```bash
# Uses default server 127.0.0.1:3478
go run ./cmd/client/main.go

# Against your own server
go run ./cmd/client/main.go <server-ip>:3478

# Against a public STUN server (no need to run your own)
go run ./cmd/client/main.go stun.l.google.com:19302
```

Example output:
```
2026/02/20 22:48:43 Public IP: 127.0.0.1, port: 52301
```
