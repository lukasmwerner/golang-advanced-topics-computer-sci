package tempconv

import (
	"flag"
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

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // No error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)

}

// CelsiusFlag defines a Celsius flag with the specified name, default value,
// and usage and returns the address of the flag variable
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
