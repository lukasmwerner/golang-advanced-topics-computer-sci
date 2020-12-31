package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x uint64 = 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		iter(&x, scanner)
	}

}

func iter(x *uint64, scanner *bufio.Scanner) {
	fmt.Printf("%v value\t%v address\r\r", *x, x)
	scanner.Scan()
	(*x)++
}
