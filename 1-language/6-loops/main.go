package main

import (
	"fmt"
)

func main() {
	/* The only loop available is the for loop.
	 *
	 * Like the if-statement there are no parenthesis around the variable
	 * instantiation, condition, nor counter expressions, and curely brackets
	 * are required on the same line.
	 */
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	/* The for loop can behave like a while loop by forgoing the variable
	 * instantiation and the counter expressions.
	 */
	 i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	/* Go has built-in support for iteration via ranges. Using the `range`
	 * operator inside a for loop with an array/slice on the RHS returns
	 * the index and the element value.
	 */
	arr := []string { "a", "d", "c", "d", }
	for index, value := range arr {

		/* fmt.Printf() behaves exactly like printf from C */
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	fmt.Println()

	/* Same principle with maps, but `range` will return the key instead
	 * of the index.
	 */
	m := map[string]int {"a" : 1, "b" : 2, "c" : 3}
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	fmt.Println()

	 /* A `while (true)` loop can be simulated using the `for` keyword, but
	  * wihtout any other information. This will run infinitely until you exit.
	  */
	fmt.Println("Infinite loop incoming...")
	for {
	}
}