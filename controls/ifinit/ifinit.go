package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	second := rand.NewSource(time.Now().UnixNano())
	randNum := rand.New(second)
	return randNum.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}
