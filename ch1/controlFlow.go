package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var heads int
	var tails int
	var selected string

	//set the random seed so that we get some randomness
	rand.Seed(time.Time.UnixNano(time.Now()))

	options := []string{"heads", "tails", "edge"}
	randIndex := rand.Intn(len(options))

	selected = options[randIndex]
	fmt.Println(selected)
	switch selected {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
	fmt.Printf("%v %v \n", heads, tails)
}
