package main

/* Lists of import packages can be written like this (not comma separated) */
import (
	"fmt"
	"time"
)

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

	/* Switch statements are also in Go. The syntax is similar to other C-like
	 * languages, the only difference is that there are no `break` statements
	 * and each `case` only supports one line.
	 */
	 	
    i := 2
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
	default:
		fmt.Println("cringe")
    }

	/* You can specify multiple cases in a sngle statement using a
	 * comma-serpated list, rather than multiple `case` statements
	 */
	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	/* There is another switch statement varient that can handles type
	 * switching , but that is reserved for the interfaces tutorial.
	 */
}