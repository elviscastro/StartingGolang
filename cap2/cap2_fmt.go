package main

import "fmt"

func main() {

	// exemplo de um print comum com caracteres interpretados
	c := "Olá\n bom \tdia, \"Elvis\"\n"

	fmt.Println(c)

	// exemplo de literal string com ``
	a := `"Olá\n bom \tdia"`

	fmt.Println(a)
	fmt.Println()

	// exemplo de uso de um Sprint, que retorna uma variavel string com o conteudo passado por parametro
	x := "oi"
	y := "bom dia"

	z := fmt.Sprint(x, " ", y)

	fmt.Println(z)

}
