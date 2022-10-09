package main

import "fmt"

func main() {
	//var aprovados map[int]string //chave int e valor string
	// maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[82893929392] = "Maria"
	aprovados[34234253455] = "Pedro"
	aprovados[34234234532] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[82893929392])
	delete(aprovados, 82893929392)
	fmt.Println(aprovados)

}
