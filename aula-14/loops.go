package main

import (
	"fmt"
	//"time"
)

func main() {

	i := 0

	for i < 10 {	
		fmt.Println("Incrementando i", i)
		i++
		time.Sleep(time.Second)
	
	}

	for j := 0; j < 10; j+=2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"João", "Maria", "José", "Pedro"}
	
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "Tomita" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome": "João",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome string
		sobrenome string
	}
	/*
	usuario2 := usuarioStruct{"João", "Silva"}

	for chave, valor := range usuario2 {
		fmt.Println(chave, valor)
	}
	*/
}