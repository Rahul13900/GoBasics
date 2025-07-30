package main

import (
	"fmt"
	"sort"
)

// sort.Interface is defined as
type Interface interface {
	Len() int           // Len is the number of elements in the collection.
	Swap(i, j int)      // Swap swaps the elements with indexes i and j.
	Less(i, j int) bool // Less reports whether the element with index i should sort before the element with index j.
}

// sort.Sort as
func Sort(data Interface) {
	// Sort sorts data in increasing order, as determined by the Less method.
	// The sort is not guaranteed to be stable.
	// It is the caller's responsibility to ensure that data implements the Interface.
	for i := 0; i < data.Len(); i++ {
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, i) {
				data.Swap(i, j)
			}
		}
	}
}

// Reverse is a function that takes a slice of integers and sorts it in reverse order.
// It does this by using the sort.Sort function and passing in a custom Less function that compares the elements in reverse order.
// User sort.Reversse which is defined as 
type reverse struct {
	Interface
}

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

func Reverse(data Interface) Interface {
	return &reverse{data}
}

func main() {
	entries := []int{5, 3, 4, 1, 2}
	// sort.Sort(entries) // this will not work as the type does not implement the interface.
	sort.Sort(sort.IntSlice(entries)) // this will work as the type implements the interface.
	fmt.Println(entries)              // [1 2 3 4 5]

	sort.Sort(sort.Reverse(sort.IntSlice(entries))) // this will work as the type implements the interface.
	fmt.Println(entries) // [5 4 3 2 1]
}
