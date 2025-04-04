package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{} 

	for _, item := range items {
		a = append(a, item[:]) // this is the reference to the original slice, not a copy
		// and since its referencing same memory each time, it will be the last value of the loop that will be stored in a for each iteration
	}
	fmt.Println(items)
	fmt.Println(a)
}
