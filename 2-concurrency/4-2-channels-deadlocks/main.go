package main

import (
	"fmt"
	"time"
)

/* Go is unique in that the built-in features make concurrency easy. The Go
 * runtime can also automatically detect when deadlocks occur (not at compile
 * time, the halting problem still hasn't solved).
 */


/* There needs to be a mechanism in the channel to tell the reciever that
 * the sender has finished sending all of it's messages.
 */
 func count(s string, c chan string) {
	for i:= 1; i < 5; i++ {
		c <- s
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	c := make(chan string)

	go count("sheep", c)

	/* This for loop will run indefinitely. The `count` goroutine has finished
	 * but the main goroutine is still waiting for a message, despite the fact
	 * that nothing will ever come to that channel. This is a deadlock
	 * condition and the program won't terminate.
	 * 
	 * Go can detect this locked state and exit the program with the proper
	 * error message.
	 */
	for {
		msg := <- c
		fmt.Println(msg)
	}
}