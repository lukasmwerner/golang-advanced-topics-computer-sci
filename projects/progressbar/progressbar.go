package main

import (
	"fmt"
	"math/rand"
	"time"
)

func updateBar(progress int, total int) {
	fmt.Printf("\r%d/%d progress ", progress, total)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	totalWorkList := 10
	for i := 1; i <= totalWorkList; i++ {
		updateBar(i, totalWorkList)
		sleeptime := time.Second * 4 / time.Duration(rand.Intn(10)+1)
		time.Sleep(sleeptime)
	}
	fmt.Println("\nFinished!")

}
