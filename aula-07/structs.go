package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}
func main() {
	fmt.Println("Aula 07 - Structs")
	
	var u usuario
	u.nome = "João"
	u.idade = 30
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua A", 10}
	
	usuario2 := usuario{"Maria", 20, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 30}
	fmt.Println(usuario3)
}