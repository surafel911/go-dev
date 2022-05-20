package main

import (
	"fmt"
)

/* Go has many design patterns for concurrency (the lagnauge is centered
 * around it numbnuts). One such pattern is the worker pool.
 *
 * Worker pools is a common pattern where a queue of work to be done
 * and  multiple concurrent workers pulling items off the queue.
 */

/* Function to calculate the fibinachi for a given number N */
func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func main() {
	num := 100
	num_goroutines := 5


	jobs := make(chan int, num)
	results := make(chan int, num)

	for i := 0; i < num_goroutines; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < num; i++ {
		jobs <- i
	}

	close (jobs)

	for i := 0; i < num; i++ {
		fmt.Println(<- results)
	}
}