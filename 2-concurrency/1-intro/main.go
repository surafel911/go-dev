package main

import (
	"fmt"
	"time"
)

func main() {
	/* As an imperitive language, Go statements run sequentially. This means
	 * that `count("sheep")` will actually never stop executing and
	 * `count("fish")` will never have the change to execute.
	 */
	count("sheep")
	count("fish")
}

func count(s string) {
	for i:= 1; true; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}