package fibonacci

// Fibonacci finds the n(th) term of the fibonacci sequence sequentially
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// GoFibonacci finds the n(th) term of the fibonacci sequence with goroutines
func GoFibonacci(n int, ch chan int) {
	if n <= 1 {
		ch <- n
		return
	}
	resultA, resultB := make(chan int), make(chan int)
	go GoFibonacci(n-1, resultA)
	go GoFibonacci(n-2, resultB)
	ch <- ((<-resultA) + (<-resultB))
}

/*
func main() {
	begin := 0
	end := 20
	fmt.Println("Fibbonacci Sequence:")
	for i := begin; i < end; i++ {
		start := time.Now()
		result := Fibonacci(i)
		fmt.Printf("\tFibonacci(%v)=%v\tTook: %v\n", i, result, time.Since(start))
	}
	for i := begin; i < end; i++ {
		ch := make(chan int)
		start := time.Now()
		go GoFibonacci(i, ch)
		result := <-ch
		fmt.Printf("\tGoFibonacci(%v)=%v\tTook: %v\n", i, result, time.Since(start))
	}

}
*/
