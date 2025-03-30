package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: echo <port_to_listen>")
	}
	listen, err := net.Listen("tcp", ":"+os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("listening on", os.Args[1])
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Error accepting connection from", conn.RemoteAddr(), ":", err)
			continue
		}
		log.Println("Connection accepted from host", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Time{}.Add(time.Duration(0)))
	buf := make([]byte, 1024)
	for {
		size, err := conn.Read(buf[0:])
		if err == io.EOF {
			log.Println("Connection closed with host", conn.RemoteAddr())
			break
		} else if err != nil {
			log.Println("Error receiving data from", conn.RemoteAddr(), ": ", err)
			break
		}
		log.Printf("Received %d bytes from %v: %s", size, conn.RemoteAddr(), string(buf[0:size]))
		if _, err := conn.Write(buf[0:size]); err != nil {
			log.Println("Error sending data to host", conn.RemoteAddr(), ": ", err)
		}
	}
}
