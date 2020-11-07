package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var port = flag.String("port", "8000", "sets the port the tcp reverb server runs on")

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
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "%s\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "%s\n", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "%s\n", strings.ToLower(shout))
}
