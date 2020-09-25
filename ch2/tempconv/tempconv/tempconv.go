package tempconv

import "fmt"

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
