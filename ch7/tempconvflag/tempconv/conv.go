package tempconv

// CToF Convert given c to fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC Convert given f to celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
