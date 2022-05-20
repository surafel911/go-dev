package main

import "fmt"

/*
 * Semi-colons *can* be used and are disscouraged by go-formatting tools.
 */

func main() {
	/* Syntax for declaring a variable */
	var x int
	var y = 6

	/* If not assigned, all variables will be assigned a "0" for their
	 * type (i.e. int=0, bool=false, string=empty string).
	 */
	fmt.Println(x)

	/* Variable used in C-like ways */
	x = 5;
	fmt.Println(x)

	/* Go can infer types using this syntax (no `auto` nor `var` like 
	 * keyword necessary). The `:=` symbol is called a
	 * `short assignment statement` and is used to delcare + assign
	 * a variable.
	 */
	sum := x + y
	fmt.Println(sum)
}
