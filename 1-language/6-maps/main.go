package main

import (
	"fmt"
)

/* In Go, `maps` are an associative data type similar to Dictionaries
 * in Java/C#, and maps in C++.
 *
 * Note that keys are not sorted nor unique.
 */

func main() {
	/* The `map` type is denoted by `map[K]T`, where K is the key and
	 * T is the type (i.e., T does not have it's own square braces).
	 *
	 * Maps can be declared with the `make` function. The `make` function
	 * is specifically used to construct and initialize built-in data types
	 * (e.g., slices, maps) and only returns that type (i.e., not a pointer
	 * to newly allocated memory).
	 * 
	 * When using `make`, the first argument is the type and the second
	 * argument is the length of the map (i.e., initial allocation size).
	 */
	var m map[string]int = make(map[string]int, 10)

	/* Maps can also be declared and assigned using the `:=` operator. Pairs
	 * are denoted by the `K : T` syntax.
	 *
	 * t := map[string]int { "price" : 10, "cost" : 5 }
	 */

	/* Printing a map shows the list of all k:t pairs inside of [] */
	fmt.Println(m)

	/* Use the subscript operator to both access and insert elements into
	 * the map.
	 */
	m["foo"] = 1
	m["bar"] = 2
	m["baz"] = 3

	fmt.Println(m)

	m["bar"] = 4
	fmt.Println(m)

	m["baz"] = 5
	fmt.Println(m["baz"])

	/* Key,Value pairs can be removed by using the `delete(map[K]T, K)`
	 * function.
	 */
	delete(m, "baz")
	fmt.Println(m)

	/* In Go, it's possible for a function/operator to return multiple values.
	 * The subcript operator returns two values: 
	 *	T: 		the value from the given key, or the zero of that type if it
	 *			does not exist in the map
	 *	bool:	a boolean as to whether a key exists in the map
	 */
	 value, found1 := m["baz"]
	 fmt.Println("[", value, ":", found1, "]")

	 /* Since the value will be zerod if the key is not in the map, skipping
	  * the assignment for that value can be done with an `_`.
	  */
	 _, found2 :=  m["bar"]
	 fmt.Println("[", found2, "]")
}