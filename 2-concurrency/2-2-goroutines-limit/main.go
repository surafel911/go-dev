package main

import (
	"fmt"
	"time"
)

func main() {
	/* Calling `count("fish")` as a goroutine as well does what's
	 * expected - making the call a goroutine and run concurrently with
	 * the main goroutine. The caviate is that the program will terminate
	 * once the main goroutine terminates, meaning that any running goroutines
	 * will terminate, even if they haven't finished their work.
	 *
	 * This results in nothing being printed. The main goroutine must be
	 * explicitly blocked to allow both goroutines to do work and prevent 
	 * the program from terminating early.
	 */
	go count("sheep")
	go count("fish")

	/* The program will close when this timer completes, even though both
	 * of the above goroutines are still capable of executing.
	 * 
	 * This is a poor solution of the problem and only exists for
	 * demostration purposes.
	 */
	time.Sleep(time.Second * 2)
}

func count(s string) {
	for i:= 1; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}