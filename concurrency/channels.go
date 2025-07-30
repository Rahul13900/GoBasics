package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		ch <- result{url, nil, time.Since(start).Round(time.Millisecond)}
		resp.Body.Close()
	}
}

func main() {
	results := make(chan result)
	list := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		res := <-results
		if res.err != nil {
			log.Printf("Error: %s: %v", res.url, res.err)
		} else {
			log.Printf("Success: %s: %v", res.url, res.latency)
		}
	}
}
