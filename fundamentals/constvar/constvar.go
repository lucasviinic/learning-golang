package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2) //sem os dois pontos (:) é no caso em que a variável já existe, (:=) estamos iniciando e atribuindo valor à variável
	// Acima, utilizamos o atalho m para referênciar ao pacote math, precisando apenas definir o atalho antes do nome do pacote
	// O Go lança erro de compilação quando uma variável é inicializada e não utilizada,
	// portanto, caso as instruções se encerracem aqui, não compilaria
	fmt.Println("A área da circunferência é", area)

	// Outras formas de declarar constantes
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Hello World!"
	fmt.Println(g, h, i)

}
