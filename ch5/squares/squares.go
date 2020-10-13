package main

import "fmt"

//kinda like a yeild in python
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	for i := 0; i <= 5; i++ {
		fmt.Printf("%v: %v\n", i, f())
	}
}
