package main

import (
	"fmt"
	"time"
)

/* Timouts are important when connecting to external resources in a
 * time dependent program. This concept is implemented with the 
 * `time.After()` method and a select statement case.
 */
func main() {
	/* Creating the channel with a buffer of 1 prevents the anonymous
	 * goroutine from hanging and allows it to close once finished.
	 */
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	
	/* The `time.After()` method returns a channel with the actual elapsed
	 * time. In this example, we only care that a message was recieved on
	 * that channel, so we simply recieve from from the returned channel in
	 * the case statement.
	 */
	timeout := time.After(time.Second * 2)

	/* In this example, the message is recieved before the timeout occurs.
	 * A result message should be printed.
	 */
	select {
	case res := <- c1:
		fmt.Println(res)
	
	/* Recieve from the timeout channel. A message will be recieved when the
	 * time has elasped.
	 */
	case <- timeout:
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	/* In this example, the message is recieved after the timeout occurs.
	 * A timout message should be printed.
	 */
	select {
	case res := <- c2:
		fmt.Println(res)
	/* For syntactic sugar you can forgo the channel variable. */
	case <- time.After(time.Second * 1):
		fmt.Println("timeout 2")
	}
}