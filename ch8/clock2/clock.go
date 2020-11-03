package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "sets the port the tcp clock server runs on")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for { // while true loop
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // smth conn aborted
		}
		go handleConn(conn) // cucurrent handler
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
