package main

import "fmt"

func fib() func() int {
	i := 0
	j := 1
	fmt.Println("i:", i, "j:", j)
	return func() int {
		// the innner function has access to the outer function's variables
		// this is a closure, the inner function "closes over" the outer function's variables
		i, j = j, i+j
		return j
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		println(f())
	}
}
