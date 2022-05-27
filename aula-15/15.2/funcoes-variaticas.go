package main

import(
	"fmt"
)


func soma(numeros ...int) int {
	total  := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int){
	total := 0

	for _,numero := range numeros{
		total += numero	
		fmt.Println(texto, total)
	}
}

func main() {
	total := soma(1,2,3,4,5)
	fmt.Println(total)

	escrever("Soma: ", 1,2,3,4,5)
}