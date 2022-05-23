package main

import (
	"fmt"
)

/* Go supports pointers and have the same syntax and semantics as in C++ */

func main() {
	i := 7
	var p *int = &i
	fmt.Println(i)
	fmt.Println(p)

	fmt.Println()

	inc(i)
	fmt.Println(i)

	incp(&i)
	fmt.Println(i)
}

func inc(x int) {
	x++
}

func incp(x *int) {
	(*x)++
}