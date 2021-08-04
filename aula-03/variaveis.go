package main

import "fmt"

func main() {
	var variavel1 string = "Variável explícita string 1"	
	variavel2 := "Variavel implícita string 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "var 3 agrupada"
		variavel4 string = "var 4 agrupada"
	)
	fmt.Println(variavel3,variavel4)

	variavel5, variavel6 := "variavel 5 linha unica", "variavel 6 linha unica"

	fmt.Println(variavel5,variavel6)

	const constante1 string = "Constante 1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5,variavel6)
}