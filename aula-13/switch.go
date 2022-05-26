package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Não é um dia da semana"
	}
}

func diaDaSemana2(numero int) string {
	var dia string
	switch{
		case numero == 1:
			dia = "Domingo"
			fallthrough
		case numero == 2:
			dia = "Segunda"
		case numero == 3:
			dia = "Terça"
		case numero == 4:
			dia = "Quarta"
		case numero == 5:
			dia = "Quinta"
		case numero == 6:
			dia = "Sexta"
		case numero == 7:
			dia = "Sábado"
		default:
			dia = "Não é um dia da semana"
	}
	return dia
}

func main() { 
	fmt.Println("Aula 13 - Switch")
	dia := diaDaSemana(4)
	fmt.Println(dia)
}