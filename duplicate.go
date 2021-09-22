package main

import "fmt"

func fibonacci() func() int {
	fib1, fib2 := 1,1
	return func() int{
		fib := fib1
		fib1, fib2 = fib2, fib1 + fib2
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
