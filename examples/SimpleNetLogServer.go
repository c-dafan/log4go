package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	port = flag.String("p", "12124", "Port number to listen on")
)

func e(err error) {
	if err != nil {
		fmt.Printf("Erroring out: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	bind, err := net.ResolveUDPAddr("udp", *port)
	//bind, err := net.ResolveUDPAddr(""，*port)
	e(err)

	// Create listener
	listener, err := net.ListenUDP("udp", bind)
	e(err)

	fmt.Printf("Listening to port %s...\n", *port)
	for {
		// read into a new buffer
		buffer := make([]byte, 1024)
		_, _, err := listener.ReadFrom(buffer)
		e(err)

		// log to standard output
		fmt.Println(string(buffer))
	}
}
