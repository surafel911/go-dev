package main

import (
	"fmt"
)

/* Sending to a channel is a blocking operation, and is a constriant to
 * working goroutines.
 */

/* To demonstrate this, look at the commented code below. The main
 * goroutine will block when sending the message because there is no
 * recieving goroutine available. Practically, imagine that the reciever
 * exists but is busy doing other things.
 *
 * What ends up happening is that the code below sending the message
 * never executes.
 *
 * func main() {
 * 	c := make(chan string)
 * 	c <- "hello"
 * 
 * 	msg := <- c
 * 	fmt.Println(msg)
 * }
 */

func main() {
	/* To resolve this issue (sending but no reciever is available), you can
	 * buffer messages until a reciever is ready to injest them.
	 *
	 * Here is the syntax for creating a channel with N buffers.
	 */
	buffer := 2
	c := make(chan string, buffer)

	/* This channel has buffer space to store two strings from the sender
	 * before blocking the sending goroutine,
	 */

	c <- "hello"
	c <- "world"

	/* If you send a message when the channel buffer is full, it will
	 * block the sender goroutine until a reciever pulls a message from
	 * the channel.
	 *
	 * Uncomment the following code to observe the behavior.
	 * 
	 * c <- "three"
	 */

	for i := 0; i < buffer; i++ {
		msg := <- c
		fmt.Println(msg)
	}
}