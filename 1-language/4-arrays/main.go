package main

import (
	"fmt"
)

func main() {

	/* Static arrays are declared with this syntax and initialized with
	 * their type's `0` value.
	 */
	var a[5] int

	/* Use subscripts to access array elements. Arrays are zero-indexed */
	a[2] = 7

	/* When given multiple values, fmt.Println will automatically space
	 * separate them in the output.
	 */
	fmt.Println("a:", a)

	/* Inline initialize a static array using the short assignment operator. */

	/* Using the `...` symbol inside the square brackets when initializing
	 * inline has Go find the length at compile time.
     * 
	 * This statment is perfectly valid as well, just some syntactic sugar:
     * b :=  [...]int{5, 4, 3, 2, 1}
     */
	b :=  [5]int{5, 4, 3, 2, 1}
	fmt.Println("b:", b)

	/* Go supports multi-dimensional arrays by adding another brace
	 * to the declaration signiture.
	 */
	var c[2][2] int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("c:", c)
}