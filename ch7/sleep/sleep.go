package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep time")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
