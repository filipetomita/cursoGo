package main

import "fmt"

var n int

func init() {
	fmt.Println("Inicializando o pacote")
	n = 10
}

func main() {
	fmt.Println("Aula 15.9 - Funções de inicialização")
	fmt.Println("valor inicializado",n)
}