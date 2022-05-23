package main

import "fmt"

func main() {
	//Aritimeticos
	soma := 2 + 3
	subtracao := 3 - 2
	divisao := 10 / 2
	multiplicacao := 2 * 3
	restoDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 5
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//Atribuição
	fmt.Println("------------------------")
	var variavel1 string = "String 1"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)

	//Relacionais 
	fmt.Println("------------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 1)

	//Lógicos
	fmt.Println("------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//Unarios
	fmt.Println("------------------------")
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 3
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)

	//Operadores ternarios
	fmt.Println("------------------------")

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
	
}