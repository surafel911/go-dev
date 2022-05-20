package main

/* Lists of import packages can be written like this (not comma separated) */
import (
	"fmt"
)

/* For further reading:
 * https://riteeksrivastava.medium.com/how-slices-internally-work-in-golang-a47fcb5d42ce
 */

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

	/* Static arrays are annoying because the length is part of the array's
	 * type. This is where slices come in.
	 * 
	 * `Slices` are an abstraction on top of arrays to make them
	 * dynamic. Slices are declared similarly to arrays, just without the 
	 * size in the square brackets.
	 *
	 * Semantically and structurally, slices are a built-in construct akin to
	 * vectors in C++.
	 */
	 var c[] int

	 /* Slices can be appended by the `append` function and passing in
	  * the original slice object and a variadic list of arguments.
	  * Append returns a *new* slice rather than modify the original slice.
	  * 
	  * Slices are backed by arrays, meaning in the background Go is
	  * creating a 
	  */
	 c = append(c, 4, 5, 6, 7, 8)
	 fmt.Println("c:", c)

	 /* Slices can be created from arrays with this syntax:
	  * a[start:end] (i.e., new slice from [start, end) or [start, end - 1]).
	  */
	  d := b[1:4]
	  fmt.Println("d:", d)

	/* Slices can be explicitly constructed using the `make([]T, len, cap)`
	 * function. 
	 *
	 * `make` allocates a zerod array and returns a slice that points to
	 * that array. Capacity is an optional argument and is (by default) equal
	 * to length when left out.
	 */
	var e[] int = make([]int, 0, 10)
	fmt.Println("e:", e)

	/* There is no default method for removing elements from a slice.
	 * Suggested method: https://stackoverflow.com/a/57213476
	 */
}