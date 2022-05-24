package main

import "fmt"



type pessoa struct {	
	nome string
	sobreNome string
	idade uint8
	altura	uint8
}

type estudante struct {
	pessoa //composição
	curso 	string
	faculdade	string
}

func main() {
	fmt.Println("Aula 08 - Herança")
	p1 := pessoa{"João", "Silva", 30, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia de Software", "UFMG"}
	fmt.Println(e1)
	fmt.Println(e1.pessoa)
	fmt.Println(e1.nome)
}