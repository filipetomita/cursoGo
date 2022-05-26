package main

import "fmt"

func main() {
	fmt.Println("Aula 12 - Estruturas de Controle")

	numero := -50

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else {
		fmt.Println("Numero menor que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Outro numero maior que 0")
	} else if outroNumero < -10{
		fmt.Println("Numero menor que -10 que é", outroNumero )
	} else {	
		fmt.Println("Numero entre 0 e -10")
	}
}