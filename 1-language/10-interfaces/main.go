package main

import (
	"fmt"
	"math"
)

/* In Go an interface is two things:
 *	1) A set of methods
 *	2) Also a type
 *
 * You define an interface using the `type *name* interface` syntax,
 * and declare a list of functions that child types are to implement.
 */
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

/* Use the `reciever type` syntax to denote inheritance and implementation
 * of an interface.
 */
func (r rect) area() float64 {
	return r.width * r.height
}

/* By implementing one interface, you are semantically required to implement
 * the rest of the functions in the interface.
 */
func (r rect) perim() float64 {
	return ((2.0 * r.width) + (2.0 * r.height))
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

/* If a variable has an interface type, then we can call the methods
 * that are implemented by the variable type.
 */
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

/* Struct types can implement multiple interfaces simply through the
 * `reciever type` syntax. The only caveat is that the you are not
 * allowed to create same name methods in two or more interfaces. Otherwise
 * the Go interpreter will panic.
 */
type shape_details interface {
	details()
}

func (r rect) details() {
	fmt.Printf("rect { width: %f, height: %f }\n", r.width, r.height)
}

func (c circle) details() {
	fmt.Printf("circle { radius: %f }\n", c.radius)
}

/* Since interfaces are fundumentally a set of methods and a type, you can
 * create something called an `empty interface`.
 *
 * An empty interface is a type with no methods. Since all types implement
 * at least zero methods, all types satisfy the empty interface. Using empty
 * interfaces is similar to using templates in C++.
 *
 * This means that any variable can be passed into the function to address
 * it's type.
 *
 * More info about empty interfaces here:
 * https://stackoverflow.com/a/23148998
 */
func what_complex_type(value interface{}) {
	/* Empty types are declared in this manner. */
	var i interface{} = value

	/* Switch cases also support type switching. The type of an empty
	 * interface can be accessed with this syntax.
	 *
	 * More info about type switching here:
	 * https://www.geeksforgeeks.org/type-switches-in-golang/
	 */
	switch t := i.(type) {
	case rect:
		fmt.Println("I'm a rect")
	case circle:
		fmt.Println("I'm a circle")
	default:
		fmt.Printf("Don't know the complex type: %T\n", t)
	}
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

	measure(r)

	fmt.Println()

    measure(c)

	fmt.Println()

	what_complex_type(r)
	what_complex_type(c)
	what_complex_type("string")
}