package main

import (
	"fmt"
)

/*
Existem duas maneiras de declaração de variavel no go, a forma explicita e implicita.

- Declaração explicita: var variavel1 string = "Variavel 1"
- Declaração implicita: variavel1 := "Variavel 1"
*/
func main() {

	var variavel1 string = "Variavel 1"
	variavel := "Variavel 2" //Inferencia de tipoq

	fmt.Println(variavel1)
	fmt.Print(variavel)

}
