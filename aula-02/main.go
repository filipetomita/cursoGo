package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá do arquivo main do novo módulo")
	auxiliar.Escrever()

	erro := checkmail.ValidadeFormat("filipe.tomita@gmailcom")
	fmt.Println(erro)
}