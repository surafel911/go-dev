package main

import (
	"fmt"
	"time"
)

/* Go has many design patterns for concurrency (the lagnauge is centered
 * around it numbnuts). One such pattern is the worker pool.
 *
 * Worker pools is a common pattern where a queue of work to be done
 * and  multiple concurrent workers pulling items off the queue.
 */

 /* Worker goroutine prints the start and finished status of the job,
  * and sleeps to simulate an expensive task.
  */
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
 }
 
func main() {
 
	const numJobs = 16
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 4; w++ {
		go worker(w, jobs, results)
	}
 
	start := time.Now()

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	/* Close the jobs channel to signal that we have no more jobs to push
	 * on the channel.
	 */
	close(jobs)
 
	/* Recieve the results on the results channel to make sure all of the
	 * goroutines have finished processing. The same effect can be achieved
	 * using waitgroups.
	 */
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	/* Notice how the elapsed time is always less than numJobs seconds (since
	 * each goroutine sleeps for one second). This is the power of
	 * concurrency.
	 */
	fmt.Println("elapsed:", time.Since(start))
 }