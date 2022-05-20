/* First line *must* be the package of the program. Binaries are required
 * to have *a* package called main.
 */
package main

/* Import statements follow the package declaration. The "fmt" package
 * is a library contains standard IO functions.
 */
import "fmt"

/* Function declaration. Note that the openning curly brace *must be on
 * the same line as the function signiture*.
 */
func main() {

	/* Semi-colons are *not* required for go to run (but are helpful ) */
	fmt.Println("Hello World");
}