package main

import (
	"fmt"
	"time"
)

func updateBar(progress int, total int) {
	fmt.Printf("\r%d/%d progress ", progress, total)
}

func main() {
	totalWorkList := 10
	for i := 0; i < totalWorkList; i++ {
		updateBar(i, totalWorkList)
		time.Sleep(time.Second / 2)
	}
	fmt.Println("\nFinished!")

}
