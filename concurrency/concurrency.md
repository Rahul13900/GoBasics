Concurrency Definations
  Basic Defination --> Concurrency in Golang means the ability to run multiple tasks independently, possibly at the same time.

  In Go, concurrency is achieved using goroutines (lightweight threads) and channels (for communication and synchronization).

Concurrency vs. Parallelism (Clarification)
  Concurrency is structuring a program to handle multiple tasks that can start, run, and complete in overlapping time periods.

  Parallelism is about running multiple tasks simultaneously, especially on multi-core CPUs.
  Go supports concurrency primarily, but can also achieve parallelism using the GOMAXPROCS setting.

We can have Concurrency with a single-core processor(think interrupt handling in the operating system), but parallelism can only happen on a multi-core processor

NOTE: Concurrency doesn't make the program faster, parallelism does.