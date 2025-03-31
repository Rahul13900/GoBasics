package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Getting the Current Time
	now := time.Now()
	fmt.Println("Current Time:", now)

	// 2. Formatting Time
	fmt.Println("YYYY-MM-DD:", now.Format("2006-01-02"))
	fmt.Println("DD-MM-YYYY:", now.Format("01-01-2006"))
	fmt.Println("Time (HH:MM:SS):", now.Format("13:04:05"))
	fmt.Println("Custom Format:", now.Format("Monday, Jan 2, 2006 at 3:04 PM"))

	// 3. Parsing Time
	timeStr := "2025-03-28 14:35:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed Time:", parsedTime)
	}

	// 4. Adding and Subtracting Time
	oneDayLater := now.Add(24 * time.Hour)
	oneWeekAgo := now.Add(-7 * 24 * time.Hour)
	fmt.Println("One Day Later:", oneDayLater)
	fmt.Println("One Week Ago:", oneWeekAgo)

	// 5. Comparing Two Dates
	t1 := time.Date(2025, 3, 25, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 3, 28, 10, 0, 0, 0, time.UTC)
	fmt.Println("t1 Before t2:", t1.Before(t2))
	fmt.Println("t1 After t2:", t1.After(t2))
	fmt.Println("t1 Equal to t2:", t1.Equal(t2))

	// 6. Time Difference (Duration Calculation)
	diff := t2.Sub(t1)
	fmt.Println("Difference in Days:", diff.Hours()/24)

	// 7. Working with Timezones
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
	} else {
		fmt.Println("Time in New York:", time.Now().In(loc))
	}

	// 8. Using Time in Goroutines
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Executed after 2 seconds")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Main function exits")

	// 9. Creating Timers and Tickers
	// Timer (Executes once after delay)
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("Waiting for timer...")
	<-timer.C
	fmt.Println("Timer expired!")

	// Ticker (Repeats at intervals)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
	}

	// 10. Measuring Execution Time
	start := time.Now()
	time.Sleep(2 * time.Second) // Simulating work
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}
