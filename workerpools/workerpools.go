package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// create worker goroutines - these will:
	// - use multiple cores
	// - pull & push onto the jobs & results queues
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		// this isn't blocking because of the channel capacity
		jobs <- i
	}

	// close the jobs channel
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// arrows show that jobs is only a receiver, and results is a sender
// compile-time error if trying to do otherwise
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// fibonatci number function
func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
