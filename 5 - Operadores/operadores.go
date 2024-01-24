package main

import "fmt"

func main() {

	//ARITMETICOS
	soma := 2 + 2
	subtracao := 1 - 2
	multiplicacao := 2 * 2
	divisao := 2 / 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// Não podemos realizar operações aritmeticas com numeros de tipos diferentes

	var numero int16 = 10
	var numero2 int32 = 10

	soma1 := numero + int16(numero2)

	fmt.Println(soma1)

	// Fim dos aritmeticos

	//ATRIBUIÇÃO

	var variavel1 string = "string"
	variavel2 := "string"

	fmt.Println(variavel1, variavel2)

	// Fim dos operadores de atribuição

	//RELACIONAIS
	/*Operadores relacionais sempre retornarão valores booleanos
	 */

	// Fim dos operadores relacionais

	//LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Fim dos operadores logicos

	// UNARIOS

	numero5 := 10
	numero5++

	numero--
	fmt.Println(numero5)

}
