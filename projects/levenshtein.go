package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 3 {
		fmt.Printf("Value: %v\n", wagnerFisher(os.Args[1], os.Args[2]))
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var args []string
		for scanner.Scan() {
			args = append(args, scanner.Text())
		}
		fmt.Printf("Value: %v\n", wagnerFisher(args[0], args[1]))
	}
}

func wagnerFisher(s string, t string) int {
	m := len(s) + 1
	n := len(t) + 1
	distances := make([][]int, n)
	for i := range distances {
		distances[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		distances[i][0] = i
	}
	for i := 1; i < m; i++ {
		distances[0][i] = i
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			var subtitutionCost int
			if iAt(s, j-1) == iAt(t, i-1) {
				subtitutionCost = 0
			} else {
				subtitutionCost = 1
			}
			value := min([]int{
				distances[i-1][j] + 1,
				distances[i][j-1] + 1,
				distances[i-1][j-1] + subtitutionCost,
			})
			distances[i][j] = value
		}
	}
	return distances[n-1][m-1]
}

func printArray(arr [][]int) {
	for _, val := range arr {
		for _, value := range val {
			fmt.Printf(" %v ", value)
		}
		fmt.Println()
	}
}

func iAt(s string, i int) string {
	runeString := []rune(s)
	selectedRune := runeString[i]
	return string(selectedRune)
}

func min(a []int) int {
	smallest := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < smallest {
			smallest = a[i]
		}
	}
	return smallest
}
