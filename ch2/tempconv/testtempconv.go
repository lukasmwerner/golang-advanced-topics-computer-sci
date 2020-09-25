package main

import (
	"fmt"
)

// Celsius  basic type
type Celsius float64

// Fahrenheit basic type
type Fahrenheit float64

const (
	// AbsoluteZeroC duh
	AbsoluteZeroC Celsius = -273.15
	// FreezingC duh
	FreezingC Celsius = 0
	// BoilingC duh
	BoilingC Celsius = 100
)

// CToF Convert given c to fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Convert given f to celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	//fmt.Printf("%g\n", boilingF-tempconv.FreezingC) // on purpose type mismatch
}
