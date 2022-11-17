package main

import "fmt"

// declarando um novo tipo chamado hotdog
type hotdog int

// declarando uma variavel tipo hotdog
var b hotdog

var c int

func main() {

	// imprimindo o tipo da variavel 'b'
	fmt.Printf("%T\n", b)

	// imprimindo o tipo da variavel 'c'
	fmt.Printf("%T", c)

}
