package main

import (
	"fmt"
)

// TODO: Definitely needs expansion (interfaces?)

/* Syntax for creating a struct. */
type person struct {
	/* Fields are simply listed here. */
	name string
	age int
}

/* Go is a procedural language, not an object-oriented one. Meaning that
 * functions that access/modify data types are no associated with types
 * (e.g., recall the `make` and `delete` functions). Return to the `C`
 * mentality.
 */
func say_hello(p person) {
	fmt.Println(p.name, "says hello!")
}

func main () {
	/* Structs can also be initialized inline. */
	p := person {name: "Surafel", age: 23}

	/* Go can list out the values of members in fmt.Println() automagically. */
	fmt.Println(p)
	fmt.Println("{", p.name, p.age, "}")

	fmt.Println()

	say_hello(p)
}