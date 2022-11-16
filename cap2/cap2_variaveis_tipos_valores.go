package main

import (
	"fmt"
)

func main() {
	x := 10
	v := 20
	z := 10.0
	y := "Olá, bom dia"
	c := x > v

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("x é maior que v? %v", c)
}
