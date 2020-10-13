package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		err := WaitForServer(url)
		if err != nil {
			fmt.Printf("Site %s is down: %v\n", url, err)
		}
		if err == nil {
			fmt.Printf("Site %v is up\n", url)
		}
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s): retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponentioal back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
