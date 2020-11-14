package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Comencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		//
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("GO!")
}
