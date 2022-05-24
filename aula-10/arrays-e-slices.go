package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Aula 10 - Arrays e Slices")

	var array1	[5]string
	array1[0] = "Primeira posição"

	fmt.Println(array1)

	array2 := [5]string{"Primeira posição", "Segunda posição", "Terceira posição", "posição alterada", "Quinta posição"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	fmt.Println("Slices")

	slice1 := []int{10,11,12,13,14}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, 15)
	fmt.Println(slice1)

	slice2 := array2[2:5]
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println("Arrays Internos")
	fmt.Println("------------------------")	

	slice3 := make([]float32,10,11)
	fmt.Println(slice3)

	slice3 = append(slice3,10)
	fmt.Println(slice3)
	slice3 = append(slice3,11)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	
	slice4 := make([]float32,5)
	fmt.Println(slice4)

	slice4 = append(slice4,10)

	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	
}