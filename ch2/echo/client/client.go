package main

import (
	"io"
	"log"
	"net"
	"os"
)

// simplified implementation of io.Copy()
func Copy(dst io.Writer, src io.Reader) (int, error) {
	var size int
	for {
		buf := make([]byte, 1024)
		sr, rerr := src.Read(buf)
		if sr > 0 {
			sw, werr := dst.Write(buf[0:sr])
			if werr != nil {
				return size, werr
			}
			size += sw
			//in case of EOF, writes to dst first
			if rerr != nil {
				break
			}
		}
	}
	return size, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: client <port_to_connect>")
	}
	conn, err := net.Dial("tcp", "127.0.0.1:"+os.Args[1])
	if err != nil {
		panic(err)
	}
	go Copy(conn, os.Stdin)
	Copy(os.Stdout, conn)
	select {}
}
