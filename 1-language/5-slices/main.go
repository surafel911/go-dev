package main

/* Lists of import packages can be written like this (not comma separated) */
import (
	"fmt"
)

/* For further reading:
 * https://riteeksrivastava.medium.com/how-slices-internally-work-in-golang-a47fcb5d42ce
 */

func main() {
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
	var a[] int

	/* Slices can be appended by the `append` function and passing in
	 * the original slice object and a variadic list of arguments.
	 * Append returns a *new* slice rather than modify the original slice.
	 * 
	 * Slices are backed by arrays, meaning in the background Go is
	 * creating a 
	 */
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println("a:", a)

	/* Slices can also be declared and initialized inline similar to arrays.
	 */
	b := []int{6, 7, 8, 9, 10}
	fmt.Println("b:", b)
	  

	/* Slices can be created from arrays with the `slice` syntax:
	  * a[start:end] (i.e., new slice from [start, end) or [start, end - 1]).
	  */
	c := a[0:2]
	fmt.Println("c:", c)

	/* Slices can be explicitly constructed using the `make([]T, len, cap)`
	 * function. 
	 *
	 * `make` allocates a zerod array and returns a slice that points to
	 * that array. Capacity is an optional argument and is (by default) equal
	 * to length when left out.
	 */
	var d[] int = make([] int, 0, 10)
	fmt.Println("d:", d)

	/* Slices are one-dimension, but can be composed to form
	 * multi-dimensional types.
	 */

	dy := 4;

	/* The only way to construct a multi-dimensional slice is with `make`. First,
	 * allocate a slice for the y-axis (i.e., how many slices to store),
	 * followed by iterating through the outer slice to allocate the inner slices.
	 */
	e := make([][]int, dy)
	for i := 0; i < dy; i++ {
		length := i + 1
		e[i] = make([]int, length)

		for j := 0; j < length; j++ {
			e[i][j] = i + j
		}
	}

	fmt.Println("e:", e)

	/* There is no default method for removing elements from a slice.
	 * Suggested method: https://stackoverflow.com/a/57213476
	 */
}