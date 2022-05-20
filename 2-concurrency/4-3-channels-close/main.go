package main

import (
	"fmt"
	"time"
)

/* To properly prevent deadlock conditions, the sender can close the
 * channel. Note that the reciever goroutine should *never* close
 * the channel since you could close the channel prematurely and cause
 * a runtime panic.
 */

func count(s string, c chan string) {
	for i:= 1; i < 5; i++ {
		c <- s
		time.Sleep(time.Millisecond * 500)
	}

	/* The sender can call the `close(chan)` function sends close
	 * signal on the channel. This informs the reciever goroutine
	 * that no more messages will be sent along this channel.
	 */
	close(c)
}

func main() {
	c := make(chan string)

	go count("sheep", c)

	/* When recieving on a channel, we can accept a second return value -
	 * a boolean which tells us if the channel is still open. We can break
	 * the for loop if the channel is no longer open.
	 *
	 * The following code shows the logic of breaking the for loop with
	 * this signal, but it is not syntactically optimal and is here for
	 * demonstration.
	 *
	 * for {
	 * 	msg, open := <- c
	 * 
	 * 	if !open {
	 * 		break
	 * 	}
	 * 
	 * 	fmt.Println(msg)
	 * }
	 */

	/* Alternatively, using syntactic sugar, we can iterate over the range
	 * of the channel using the following syntax. Using the range statement
	 * will keep recieving messages until the the channel has been closed.
	 */
	for msg := range c {
		fmt.Println(msg)
	}
}