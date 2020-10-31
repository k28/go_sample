package main

import "fmt"

func format_string() {
	str := fmt.Sprintf("make format string [%d]", 7)
	fmt.Printf("%s\n", str)
}

func main() {
	format_string()
}
