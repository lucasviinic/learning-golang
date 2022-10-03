package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i //Pegando endereço da variável
	*p++   //desreferenciando (pegando valor)
	i++

	//Go não tem aritmética de ponteiros
	//p++

	fmt.Println(p, *p, i, &i)
}
