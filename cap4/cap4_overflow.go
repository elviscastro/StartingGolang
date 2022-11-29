package main

import "fmt"

func main() {
	var i uint16 // tem limite de 0 a 65535
	i = 65535
	fmt.Println(i)
	// i = 65536  // simulando estouro de memoria (Overflow)
	// fmt.Println(i)
	i++
	fmt.Println(i)
}
