package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)        // por valor, não causa efeito colateral
	fmt.Println(n) // Out: 1

	inc2(&n)       // por referência, causa efeito colateral
	fmt.Println(n) // Out: 2

	// Obs: evitar a passagem por referência sempre que possível, e priorizar a passagem por valor
}
