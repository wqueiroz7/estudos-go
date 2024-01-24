package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Todos os items do array precisam ser obrigatoriamente do mesmo tipo
	var array1 [5]int
	//Sempre iniciando da posição 0, precisamos obrigatoriamente declarar o tamanho
	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// Fixa o tamanho do array de acordo com a quantidade de items(valores) que eu passei.
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slice não é um array, mas é como se fosse uma fatia de um array.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//Adicionando valores no slice
	slice = append(slice, 25)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 22
	fmt.Println(array2)
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	//Cria um array de 15 posições e retorna um slice que pega as 10 primeiras posiçõs do array.
	// Quando o Go ve que o tamanho do slice vai estourar, ele dobra a capacidade
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
