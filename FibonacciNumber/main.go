package main

import "fmt"

type fibonacci struct {
	n1 uint64
	n2 uint64
}

func (fib *fibonacci) fibonacciNumber(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	fib.n1 = 0
	fib.n2 = 1
	for i := 3; i <= n; i++ {
		if i&1 == 1 {
			fib.n1 = fib.n1 + fib.n2
		} else {
			fib.n2 = fib.n1 + fib.n2
		}
	}

	if n&1 == 1 {
		return fib.n1
	} else {
		return fib.n2
	}
}

func main() {
	var a fibonacci
	fmt.Printf("Fibonacci 8 = %d \n", a.fibonacciNumber(8))
}
