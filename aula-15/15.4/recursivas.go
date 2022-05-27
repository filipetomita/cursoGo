package main

import (
	"fmt"
)

func fibonacci(posicao uint) uint{
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao - 1) + fibonacci(posicao - 2)
}


func main() {
	fmt.Println("Aula 15.4 - Recursivas")
	//Recursividade Fibonacci	
	posicao := uint(20)

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(posicao))
}