package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //switch true, o valor true pode ser passado mas tbm pode ser omitido
	case t.Hour() < 2:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite.")
	}
}
