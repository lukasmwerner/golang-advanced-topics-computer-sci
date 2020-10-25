package fibonacci

import (
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	param := 10
	for i := 0; i < b.N; i++ {
		Fibonacci(param)
	}
}

func BenchmarkGoFibonacci(b *testing.B) {
	param := 10
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go GoFibonacci(param, ch)
		<-ch
	}
}

func TestFibonacci(t *testing.T) {
	result := Fibonacci(9)
	if result != 34 {
		t.Errorf("Fibonacci was wrong, got: %d, want: %d", result, 34)
	}
}

func TestGoFibonacci(t *testing.T) {
	ch := make(chan int)
	go GoFibonacci(9, ch)
	result := <-ch
	if result != 34 {
		t.Errorf("GoFibonacci was wrong, got: %d, want: %d", result, 34)
	}
}
