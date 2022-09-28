// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes e
para usá-los é necessário declarar um ou vários imports
*/
import "fmt" //Pois dentro desse pacote tem IOs formatados

// A função main é a porta de entrada de um programa Go
func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")

	/*
		Sobre comentários...

		1) Priorize código legível e faça comentários que agregam valor!
		2) Evitar comentários óbvios
		3) Durante o curso abusar de comentários
	*/
}
