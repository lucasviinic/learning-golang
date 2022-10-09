package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s, len(s), cap(s))

	s = make([]int, 10, 20)        //O slice tem 10 posições, porém o array interno criado possui 20 posições
	fmt.Println(s, len(s), cap(s)) //cap retorna a capacidade máxima do array que o slice referencia

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
