package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aula 11 - Maps")

	usuario := map[string]string{
		"nome": "João",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro":"João", 
			"ultimo":"Silva",
		},
		"curso":{
			"nome":"Engenharia de Software",
			"faculdade":"UFMG",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome":"Aries",
	}
	fmt.Println(usuario2)
}