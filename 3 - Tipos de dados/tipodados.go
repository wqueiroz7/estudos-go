package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		- O tipo inteiro no go tem diversos tamanhos
		int8, int16, int32, int64
		O que diferencia é somente o tamanho de byte que o numero pode ter(tamanho)
		Quando não especificamos o tamanho ele vai usar a arquitura do meu computador para setar o Int ex; 32 ou 64 bytes
		O tipo uint quer dizer um int sem sinal

		Quando não iniciamos uma variavel, o go atribui um valor a ela implicitamente.
		O tratamento de erro dentro do Go é muito diferente do convencional, então por isso o erro dentro do Go é um tipo.

	*/
	// INT
	var numero int16 = 100
	numero2 := 10000000000000
	var numero3 uint32 = 100
	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FLOAT
	var numeroreal1 float32 = 123.45
	var numeroreal2 float64 = 123.45
	numeroReal3 := 123.456 //INFERENCIA pega o tipo da arquitetura
	fmt.Println(numeroreal1)
	fmt.Println(numeroreal2)
	fmt.Println(numeroReal3)

	// STRING

	str2 := "Texto"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char)

	// BOOL

	var booleano1 bool
	fmt.Println(booleano1)

	// ERROR

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
