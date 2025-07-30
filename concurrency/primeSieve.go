package main

import "fmt"

func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}
	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)
	go generator(limit, ch)

	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		next := make(chan int)
		go filter(ch, next, prime)
		ch = next
		fmt.Print(prime, " ")
	}
}

func main() {
	sieve(100)
}
