package main

import (
	"fmt"
)

/* There are times when a program will panic but needs to be able to recover.
 * For example, handling a client request causes a critical error, but the
 * server still needs to serve other client requests. It's necessary to recover
 * from the error.
 * 
 * In principle this is similar to exception handling in other languages.
 */

 /* Function to panic. */
func my_panic() {
	panic("problems")
}

func main() {
	/* Recovering is only useful inside a deferred function. The `recover()`
	 * keyword recieves a `recover` type. If panic was called, then the if
	 * statement will return true and will execute the recovery code.
	 */
	defer func() {
		if r:= recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	/* It's also possible to use the type assertion syntax to check if the
	 * error recieved is of a specific type.
	 *
	 * if r, ok := recover().(error); ok {
     * 	fmt.Println("Recovered in f", r)
	 * }
	 */

	my_panic()

	fmt.Println("Did not recover from panic.")
}