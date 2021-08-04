package main

import (
	"fmt"
	"errors"
)

func main() {
	var numero int64 = -100000000
	fmt.Println(numero)
	
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123654
	fmt.Println(numero3)

	//INT8 = BYTE
	var numero4 byte = 128
	fmt.Println(numero4)

	var numeroReal1 float32 = 321.12
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 987654.12
	fmt.Println(numeroReal2)

	numeroReal3 := 321.123
	fmt.Println(numeroReal3)

	//String 
	var str string = "Texto plano"
	fmt.Println(str)

	str2 := "Texto Plano 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var inteiro int64
	fmt.Println(inteiro)

	//Bool

	var booleano1 bool = false
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	var erro error = errors.New("erro Interno")
	fmt.Println(erro)
}
