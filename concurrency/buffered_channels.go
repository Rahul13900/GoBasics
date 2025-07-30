package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	semp := make(chan int, 3)
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			semp <- 1 // Acquire a semaphore slot
			defer func() { <-semp }()
			fmt.Printf("Starting task %d\n", taskID)
			time.Sleep(1 * time.Second)
			fmt.Printf("Finished task %d\n", taskID)
		}(i)
	}
	wg.Wait()
}
