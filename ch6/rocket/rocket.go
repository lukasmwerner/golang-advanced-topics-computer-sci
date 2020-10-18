package main

import (
	"fmt"
	"time"
)

type Rocket struct {
	acceleration float64
	state        string
}

func (r *Rocket) Launch() {
	r.state = "liftoff"
}

func main() {
	r := new(Rocket)
	r.state = "initalization"
	fmt.Printf("State: %v\n", r.state)
	liftoff := time.Now().Add(time.Second * 10)
	time.AfterFunc(10*time.Second, r.Launch) // say you are gonna run the func in 10 secs
	countdown := -10
	for time.Now().Before(liftoff.Add(time.Second * 5)) {
		if countdown < 0 {
			fmt.Printf("T%v\tState: %v\n", countdown, r.state)
		} else {
			fmt.Printf("T+%v\tState: %v\n", countdown, r.state)
		}

		time.Sleep(time.Second)
		countdown++
	}
	fmt.Printf("After (~14s) State: %v\n", r.state)
}
