package main

import "fmt"

func main() {
	var x int

	/* Standard I/O functions behave similarly to C standard I/O functions. */

	/* Similar to C++, Go makes uses pointers and references using
	 * `*` and `&` syntax. Go has a garbage collector, but still opts
	 * to let the user access memory.
	 */
	fmt.Scanln(&x)

	/* Here is an if-else statement. Syntactically the main differences b/w
	 * Go and other C-like languages are:
	 * 		1) no parenthesis around the condition expression
	 *		2) Curly brackets on the same line are required for the scope
	 */
	  	
	if x == 5 {
		fmt.Println("Equal than 5!")
	} else if x > 5 {
		fmt.Println("Greater than 5!")
	} else {
		fmt.Println("Less than 5!")
	}
}