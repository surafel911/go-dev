package main

import (
	"fmt"
	"time"
)

/* If there is a goroutine that recieves from multiple channels, it's
 * possible that one channel is ready to send data earlier than another
 * channel. But that synchronously recieving from the second channel
 * would block the recieveing goroutine from recieving a message from the
 * faster goroutine.
 *
 * This is an illustration of that problem.
 */

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	/* This anonymous goroutine sends a message every half second. */
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	/* This anonymous goroutine sends a message every two seconds. */
	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	/* Our main goroutine recieves messages from both channels. The two
	 * Println() statements execute one-after-another, despite that
	 * `fmt.Println(<- c1)` is ready to print more often than
	 * `fmt.Println(<- c2)`.
	 *
	 * This is because `fmt.Println(<- c2)` blocks the main goroutine until 
	 * c2 has a message to give to the reciever, blocking data thats ready
	 * and available from c1 from being processed.
	 */
	for {
		fmt.Println(<- c1)
		fmt.Println(<- c2)
	}
}