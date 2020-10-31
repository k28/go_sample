package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// go f(x, y, z) で新しいgoroutineが実行される
	go say("world")
	say("Hello")
}
