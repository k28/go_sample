package main

import "fmt"

// struct is collection of fields
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0, Y:0
	vp = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	// pointers to struct
	p := &v
    // (*p).XでなくてもOK
	p.X = 24
	fmt.Println(v)

	fmt.Println(v1, vp, v2, v3)
}
