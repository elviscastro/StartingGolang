package main

import (
	"fmt"
)

type tipo int

var x tipo

type tipo2 tipo

var y tipo2

func main() {
	fmt.Printf("%v\n%T\n", x, x)
	x = 42
	fmt.Printf("%v\n", x)
	y = tipo2(x)
	fmt.Printf("%v\n%T\n", y, y)
}
