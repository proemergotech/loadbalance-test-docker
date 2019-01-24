package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var msg string

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Panic("specify string to echo")
	}

	msg = args[1] + "\n"

	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Panicf("Error listening: %+v", err.Error())
	}

	log.Printf("Listening on: %v with message: %v", l.Addr().String(), msg)

	// Close the listener when the application closes.
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()

	log.Printf("Accepted connection from: %v", conn.RemoteAddr().String())
	defer log.Printf("Closed connection from: %v", conn.RemoteAddr().String())

	go func() {
		for {
			_, err := conn.Write([]byte(msg))
			if err != nil {
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
	}

}
