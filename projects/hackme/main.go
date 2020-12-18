package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x uint16 = 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		iter(&x, scanner)
	}

}

func iter(x *uint16, scanner *bufio.Scanner) {
	fmt.Printf("%v value\t%v address\r\r", *x, x)
	scanner.Scan()
	(*x)++
}
