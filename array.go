package main

import "fmt"

func main() {
	// The type [n]T is an array of n values of type T.
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// The type []T slice with elements of type T
	// Thes selects a half-open range which includes then first one, but excludes the last one.
	var s []int = primes[1:4] // [3 5 7]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changeing the elements of a slice modifies the corresponding elements of its underlying array/
	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)
}
