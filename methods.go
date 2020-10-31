package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about Type %T!\n", v)
	}
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)

	p := Vertex{5, 7}

	fmt.Println(v, p)

	do(21)
	do("hello")
	do(true)
}
