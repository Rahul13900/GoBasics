Closures is when a function inside another function "closes over" one or more local variable of the outer function.

The Closures take refernce to the variable and not the copy, so the address remains same for a particular variable.


Examples of Closures at Industry level:

Closures can be used to implement rate limiters for API requests, ensuring that clients do not overwhelm a service.

Example: Rate Limiter using Closures

package main

import (
	"fmt"
	"time"
)

// RateLimiter returns a closure that limits function execution
func RateLimiter(limit int, interval time.Duration) func() bool {
	count := 0
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			count = 0 // Reset the count after the interval
		}
	}()

	return func() bool {
		if count < limit {
			count++
			return true
		}
		return false
	}
}

func main() {
	allowRequest := RateLimiter(3, time.Second) // Allow 3 requests per second

	for i := 0; i < 5; i++ {
		if allowRequest() {
			fmt.Println("Request allowed:", i+1)
		} else {
			fmt.Println("Request denied:", i+1)
		}
		time.Sleep(300 * time.Millisecond) // Simulating API calls
	}
}

Why Closures are Powerful in Go?
✅ Encapsulation → Keep state without exposing global variables.
✅ Dynamic Behavior → Modify logic at runtime without changing function signatures.
✅ Concurrency-Friendly → Use closures in Goroutines, middleware, and rate limiters.