package main

import "fmt"

func array() {
	// The type [n]T is an array of n values of type T.
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Loop
	for i, v := range primes {
		fmt.Println(i, v)
	}
	// if not need index
	// for _, v := range primes

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
	// Changeing the elements of a slice modifies the corresponding elements of its underlying array.
	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)
}

func arraySlice() {
	fmt.Println("--------- array slice ---------")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[1:4]
	printSlice(s)

	s = s[:2]
	printSlice(s)

	q := s[1:]
	printSlice(q)
	printSlice(s)
}

func sliceAppend() {
	fmt.Println("--------- slice append ---------")

	var s []int
	printSlice(s)

	// append works on nil slice
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func rangeSample() {
	fmt.Println("--------- range sample ---------")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// ranging over a slice, two values are returnd for each iteration.
	// index and value(copy of the element)
	// can skip the index or value by assigning to _.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func main() {
	array()
	arraySlice()
	sliceAppend()
	rangeSample()
}
