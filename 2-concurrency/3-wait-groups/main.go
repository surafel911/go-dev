package main

import (
	"fmt"
	"time"

	/* Wait groups need the sync library */
	"sync"
)

/* A wait group is a data structure to track the number of goroutines
 * that the inline goroutine should wait on.
 */

func main() {
	/* Syntax for declaring a wait group variable. This data type is in the 
	 * `sync` package and currently uses dot-notation to access the type in
	 * the package.
	 */
	var wg sync.WaitGroup

	/* The `Add(int)` function accepts an integer that adds to a counter. This
	 * counter allows the calling goroutine wait until the counter returns
	 * to 0 after the child goroutines called `Done()`.
	 */
	wg.Add(2)

	/* `Done()` decrements the counter tracking the number of goroutines
	 * to wait on.
	 */

	/* Calling `wg.Done()` is often not the responsibility of the goroutine.
	 * Rather then passing a pointer to the goroutine, it's better to call a
	 * goroutine of an anonymous function. The anonymous goroutine should
	 * call the intended function and `wg.Done()`.
	 */
	go func() {
		count("sheep")
		wg.Done()
	}()

	go func() {
		count("fish")

		/* The `defer` keyword defers the execution of a function until the
		 * surrounding function returns. The arguments are evaluated
		 * immediately, but execution does not start until the surrounding
		 * function returns.
		 *
		 * In this context, defer defers execution until the end of the
		 * anonymous function, and has the same behavior as the above anonymous
		 * function.
		 */
		defer wg.Done()
	}()
	
	/* `Wait()` blocks the calling goroutine until the the counter is 0. */
	wg.Wait()
}

func count(s string) {
	for i:= 1; i < 5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}