package main

import "fmt"

func fibonacci() func() int {
	// initial value
	var n = 0
	var b = 0
	return func() int {
		if n == 0 && b == 0 {
			n = 0
			b = 1
			return 0
		} else {
			tmp := b + n
			b = n
			n = tmp
			return tmp
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
