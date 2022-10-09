package main

import "fmt"

// Em Go, uma função também é um clousure
// Ou seja, a função tem ciência do local onde ela foi definida, ela trás as informações de onde ela foi definida
func clousure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := clousure()
	imprimeX()
}

//Out:
//20
//10
