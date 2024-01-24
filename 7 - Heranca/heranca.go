/*
A linguagem Go não possui herança propriamente dito, como as linguagens de OO, mas possui uma pseudo herança conforme o codigo abaixo
*/

package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	curso     string
	faculdade string
	pessoa
}

func main() {
	fmt.Println()

	p1 := pessoa{nome: "Jony", idade: 20, altura: 178}

	e1 := estudante{"Engenhari", "Usp", p1}

	fmt.Println(e1.nome, e1.idade, e1.altura, e1.curso)
}
