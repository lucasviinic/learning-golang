package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(2.5))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Google"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
