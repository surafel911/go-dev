package main

import (
	"fmt"
	"time"
)

/* Go implements a Communicating sequential Processes (CSP) model
 * via channels.
 *
 * Channels are a data structure that take a type and provide an
 * opportunity for two gorountines to synchronize and send data. Channels
 * are bidirectional by default *but should be used unidirectionally*.
 */

func main() {
	/* Syntax for creating a channel. A channel works with only one type.
	 * This channel can only send/recieve string values.
	 */
	c := make(chan string)

	/* Call a goroutine that accepts a channel. Note how this passed channel
	 * does not need to be a pointer to function.
	 */
	go count("sheep", c)

	/* Syntax for recieving a message from the channel. The calling (in this
	 * case main) goroutine will halt until the sender sends a message on
	 * this channel.
	 *
	 * When the other goroutine executes line to send a message, that
	 * goroutine will half and be a moment of synchronization to transfer
	 * data.
	 */
	msg := <- c

	/* Since there was only one recieve statement with the channel, the main
	 * goroutine will continue executing right after assigning `msg`.
	 */
	fmt.Println(msg)

	/* Only one "sheep" is printing to standard output because the main
	* goroutine offers only one opportunity to synchronize and accept data.
	*/
}

/* For CSP-style messaging, the goroutine must accept a channel of the
 * same type.
 */
func count(s string, c chan string) {
	for i:= 1; i < 5; i++ {
		/* This syntax sends a message along the channel. This call blocks
		 * the goroutine until a reciever is ready to reieve the message.
		 * The message is syntactically required to be of the same time.
		 *
		 * With this implementation, only the first message gets sent. Since
		 * there aren't multiple recieve statements with this channel, this
		 * will block until the main goroutine returns, casuing this goroutine
		 * to stop executing prematurely.
		 */
		c <- s
		time.Sleep(time.Millisecond * 500)
	}
}