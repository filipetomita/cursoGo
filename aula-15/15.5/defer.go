package main

import (
	"fmt"
)

func funcao1 (){
	fmt.Println("Função um - Aula 15.5 - Defer")
}

func funcao2 (){
	fmt.Println("Função dois - Aula 15.5 - Defer")	
}


func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada")
	fmt.Println("Verificando se o aluno está aprovado...")
	media := (n1 + n2) / 2
	if media >= 7 {
		
		return true
	}
	return false
}


func main() {
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(6,7))
}