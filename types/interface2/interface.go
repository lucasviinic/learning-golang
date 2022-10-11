package main

import "fmt"

type esportivo interface {
	ligarTubo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// Para que ferrari possa ser usada a partir da interface esportivo,
// vamo criar uma função ligarTubo associado a struct de ferrari
// por meio de receiver

func (f *ferrari) ligarTubo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTubo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTubo()

	fmt.Println(carro1, carro2)
}
