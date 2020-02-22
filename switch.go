package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go run on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd, windows...
		fmt.Println("%s.\n", os)
	}

	// switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Gool evening.")
	}
}
