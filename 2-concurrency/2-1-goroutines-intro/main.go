package main

import (
	"fmt"
	"time"
)

func main() {
	/* By adding the `go` keyword to a function call, it makes the function
	 * run concurrently. This is called a `goroutine`. Even the main thread is
	 * executed on a goroutine. A Goroutine is an independently executing
	 * function with it's own stack.
	 * 
	 * Now this program has two goroutines: main and count("sheep"). It
	 * should print sheeps and fishes. Goroutines are cheap to create/destroy
	 * but realize they are limited by the hardware (i.e., # of logical
	 * threads).
	 *
	 * Goroutines *are not threads*. Goroutines can be multiplexed into N>=1
	 * threads running in a program. However, thinking goroutines are as
	 * cheap as threads isn't far off.
	 */
	go count("sheep")
	go count("fish")
}

func count(s string) {
	for i:= 1; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}