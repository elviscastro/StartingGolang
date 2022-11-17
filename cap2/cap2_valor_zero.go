package main

import (
	"fmt"
)

var x int
var y float64
var z bool
var s string

func main() {

	// imprimindo cada tipo com seus 'zeros'

	fmt.Printf("x: %v (%T)\n", x, x)
	fmt.Printf("y: %.1f (%T)\n", y, y)
	fmt.Printf("z: %v (%T)\n", z, z)
	fmt.Printf("s: %v (%T)", s, s)

}
