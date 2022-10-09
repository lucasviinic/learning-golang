package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel S": 1500,
			"Guns K":    75000.32,
		},
		"J": {
			"José G": 5250.63,
		},
		"C": {
			"Carl L": 6388.99,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
