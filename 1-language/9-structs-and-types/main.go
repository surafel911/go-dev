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

/* Go does not have inheritence, rather it supports composition of
 * types. To compose complex types, treat it like build-in types or
 * simply the type name where it will act as the member name and type.
 */
type person_info struct {
	person
	id int
}

/* Go supports member functions (i.e., methods). Methods are declared and
 * defined outside of the struct. Rather than using `this`, define a
 * `reciever type` using this syntax.
 */
func (p person) say_hello() {
	fmt.Println(p.name, "says hello!")
}

func main () {
	/* Structs can also be initialized inline. */
	p := person {name: "Surafel", age: 23}
	info := person_info { person: p, id: 1234}

	/* Go can list out the values of members in fmt.Println() automagically. */
	fmt.Println(p)
	fmt.Println("{", p.name, p.age, "}")

	fmt.Println()

	fmt.Println(info)

	fmt.Println()

	p.say_hello()
}