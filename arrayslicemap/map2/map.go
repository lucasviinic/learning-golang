package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José J":    11340.45,
		"Gabriel F": 34000.5,
		"Pedro K":   1250.0,
	}

	funcsESalarios["Gabriel F"] = 1350.0
	delete(funcsESalarios, "Inexistente") //Não gera erro

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
