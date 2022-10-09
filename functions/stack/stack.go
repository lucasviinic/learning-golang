package main

import "runtime/debug"

func f3() {
	debug.PrintStack() // Imprime pilha de execução de um rpograma no momento em que a funão é chamada
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
