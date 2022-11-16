package main

import (
	"fmt"
)

// Aqui ocorre uma declaração global, acessivel por qualquer codeblock (função criada delimitada entre chaves '{}')
var y = 10

func main() {
	z := 20

	qualquercoisa(z)
}

func qualquercoisa(x int) {
	fmt.Printf("Valor 'var' (global) %v\n", y)
	fmt.Printf("Valor declarado no 'main' %v", x)
}
