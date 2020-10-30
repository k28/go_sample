package main

import "fmt"

func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// delete value with key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 存在チェック
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	mutatingMaps()
}
