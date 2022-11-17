package main

import "fmt"

// declarando um novo tipo chamado hotdog
type hotdog int

// declarando uma variavel tipo hotdog
var b hotdog = 12

var c int = 10

func main() {

	// imprimindo o tipo da variavel 'b'
	fmt.Printf("b: %v %T\n", b, b)

	// imprimindo o tipo da variavel 'c'
	fmt.Printf("c: %v %T\n", c, c)

	// essa atribuição dará erro
	// c = b

	// utilizando conversão, funcionará a atribuição anterior
	c = int(b)

	// imprimindo o tipo da variavel 'c'
	fmt.Printf("c: %v %T\n", c, c)
	c := 10

	// utilizando conversão de um tipo criado
	b = hotdog(c)

	// imprimindo o tipo da variavel 'b'
	fmt.Printf("b: %v %T\n", b, b)
}
