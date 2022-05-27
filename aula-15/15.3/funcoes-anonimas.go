package main

import (
	"fmt"
)

func main() {
	func (texto string){
		fmt.Println("Olá mundo!", texto)
	}("Go")

	retorno := func (texto string) string{
		return fmt.Sprintf("Olá mundo! %s Aula %d", texto, 15)
	}("GoLang")

	fmt.Println(retorno)
}