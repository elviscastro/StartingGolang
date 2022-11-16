package main

import (
	"fmt"
)

// declaração de variavel com tipo especificado
var x int = 10

func main() {
	// provocando erro ao atribuir valor de tipo 'float' em variavel 'int'
	// x = 20.5
	// retornando 'cannot use 20.5 (untyped float constant) as int value in assignment (truncated)'
	fmt.Println(x)
	x := 20
	fmt.Println(x)
}
