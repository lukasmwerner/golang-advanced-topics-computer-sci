package main

import (
	"fmt"
	"time"
)

const classic = `-\|/`
const updown = `▁▃▄▅▆▇█▇▆▅▄▃`
const moonphases = `🌑🌒🌓🌔🌕🌖🌗🌘`
const boxPhase = ` ░▒▓█▓▒░`
const npm = `⣷⣯⣟⡿⢿⣻⣽⣾`

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow operation
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range boxPhase {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
