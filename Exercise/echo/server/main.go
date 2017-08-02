package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	ln, err := net.Listen("tcp", ":10080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "tcp died: %s\n", err)
		os.Exit(1)
	}

	go func() {
		fmt.Println("Waiting client...")
		for {
			conn, err := ln.Accept()
			if err != nil {
				// handle error
				fmt.Printf("Failed to shake hand with client: %s", err)
				continue
			}
			go handleConnection(conn)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	<-c

	log.Fatal(ln.Close())
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 256)
	n, err := conn.Read(b)
	if err != nil {
		fmt.Printf("Failed to receive message: %s\n", err)
		return
	}
	if n <= 0 {
		fmt.Println("Received zero length message")
	}
	msg := b[0:n]
	fmt.Printf("Received: %s", string(msg))
	var echo string
	if strings.Contains(strings.ToLower(string(msg)), "hello") {
		echo = "World!\n"
	} else {
		echo = fmt.Sprintf("You sent: %s\n", string(msg))
	}
	written, err := conn.Write([]byte(echo))
	if err != nil {
		fmt.Printf("Failed to send response: %s, %v", err, written)
	}
}
