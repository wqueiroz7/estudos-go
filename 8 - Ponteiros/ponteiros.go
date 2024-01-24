/*
Ponteiro é uma variavel, que vai salvar nela, não um valor necessariamente, mas o endereço de memoria. Ele é uma reference de memoria.

*/

package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1 += 15

	fmt.Println(variavel1, variavel2)

	var variavel3 int = 100
	//Ponteiro vazio retorna Nil
	// O retorno é o endereço de memoria dentro do computador
	var ponteiro *int = &variavel3
	fmt.Println(variavel3, ponteiro)

	//Para ver o valor do ponteiro - desreferenciação
	fmt.Println(*ponteiro)
}
