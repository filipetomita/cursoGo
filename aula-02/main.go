package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá do arquivo main do novo módulo")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("filipe.tomita@gmail.com")
	fmt.Println(erro)
}