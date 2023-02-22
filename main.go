package main

import (
	"fmt"
	"net"
	"os"
)

const DEFAULT_ADDRESS = "0.0.0.0:8080"

func main() {

	address := os.Getenv("LISTEN_ADDR")
	if address == "" {
		address = DEFAULT_ADDRESS
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Fprintf(os.Stdout, "TCP server listening on %s \n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			for {
				buffer := make([]byte, 1024)
				len, err := conn.Read(buffer)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error occured while reading data from the socket %#v \n", err)
					return
				}

				fmt.Fprintf(os.Stdout, "message received: %s \n", string(buffer[:len]))
				conn.Write([]byte("message received \n"))
			}
		}(conn)
	}
}
