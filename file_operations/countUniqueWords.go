package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println("Number of Unique words", len(words))

	// define a struct to have key , value pair
	type kv struct {
		key   string
		value int
	}
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}
	// sort to slice in descending order based on value(i.e. occurance of each word)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})

	for _, s := range ss {
		fmt.Println(s.key, "apppears", s.value, "times")
	}

}
