package main

import (
	"fmt"
)

type usuario struct {
	nome string
	idade	uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando usuario %s com idade %d\n",u.nome,u.idade)
}

func (u usuario) eMaiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
	fmt.Printf("Usuário %s fez %d anos\n",u.nome,u.idade)
}

func main() {
	usuario1 := usuario{"Filipe",38}
	usuario1.salvar()
	fmt.Println(usuario1)

	usuario2 := usuario{"João",17}
	usuario2.salvar()
	fmt.Println(usuario2)
	fmt.Println(usuario2.eMaiorDeIdade())
	usuario2.fazerAniversario()
	fmt.Println(usuario2.eMaiorDeIdade())
}