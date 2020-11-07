package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn) // Note: we are ignoring errs
		log.Println("done")
		done <- 1

	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for stdout io printer conn to close.
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
