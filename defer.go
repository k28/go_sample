package main

import "fmt"

func main() {
	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding function returns.
	defer fmt.Println("world")

	for i := 0; i < 10; i++ {
		// deferred function calls are pushed onto a stack.
		// deferred calls are executed in last-in-first-out order
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}
