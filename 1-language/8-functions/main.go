package main

import (
	"errors"
	"fmt"
	"math"
)

/* The syntax for creating a function is slightly different. The `func` keyword
 * comes before the function name. As before, the type comes after the
 * variable name, and the return type is the last thing in the
 * function signiture.
 *
 * Functions in Go do not need forward declarations. This definition could be
 * below main and it would work perfectly fine.
 *
 * Void functions (i.e., functions that do not return anything) simply omit a
 * return type. The `void` keyword does not exist in Go.
 */	
func sum(x int, y int) int {
	return x + y
}

/*
 * Functions can have multiple return values. Denoted as a comma-separated
 * list of types wrapped in parenthesis. All return statements need to return
 * all of the demanded types.
 *
 * Most commonly used for `error`.
 */
func sqrt(x float64) (float64, error) {
	if x < 0 {
		/* Errors in Go are unique. They are not exceptions (i.e., no stack 
		 * tracing, etc.), and they do not halt execution. `Error` is an
		 * interface type that handles error messages and nested errors
		 * (e.g., a microservice error wrapping a DBMS error).
		 */
		return 0, errors.New("Undefined for negative numbers.")
	}

	/* `nil` is essentially null for many types, including error types.
	 * Used to signal that no errors happened during this function call.
	 */
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(sum(4, 5))

	/* When a function has two return variables, you *should* handle them
	 * either with `_` or by assigning them to a variable. Syntactically
	 * you are capable of skipping assignment and error checking, but it's
	 * partially incentivized by the compiler.
	 * 
	 * These are all valid call variations:
	 * 		1) sqrt(-9)
	 * 		2) _, err := sqrt(-9)
	 * 		3) result, _ := sqrt(-9)
	 *		4) result, err := sqrt(-9)
	 */
	result, err := sqrt(-9)

	/* If err is not handled, then nothing will happen to execution.
	 * It's a *very* bad decision to *not* handle errors, since result
	 * will either be the zero for the return type or some unspecified value.
	 */
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	/* Take careful note that execution did not halt and this fmt.PrintLn
	 * still executed. `error` is used to signal and have the programmer 
	 * safely recover, not necessarily to halt execution.
	 */
	fmt.Println("reached")
}