package main

import "fmt"

func main() {
	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		i2 := i // create a new variable to capture the value of i at that time
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}
	for i := 0; i < 4; i++ {
		// the value of i is captured by reference, not by value
		// so when the function is called, it will print the value of i at that time
		// which is 4 in this case, because the loop has finished so it gives same value and address as well
		s[i]()


		// to resolve this issue, we can use a new variable to capture the value of i at that time
	}
}
