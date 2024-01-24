/*
Functions tb é um tipo de dado dentro do go

- Funcao é uma serie de passos que serão executadas e irão fazer o que essas instruções estão mandando.
- Funcao pode retornar um valor e ter parametros.
- Funcao pode ter mais de um retorno
-  Quando chamamos a função por ele ter dois retornos precisamos setar duas variaveis pra inferir elas
- Se não quisermos usar 2 variaveis podemos setar um dos retornos com _ e ignorar um dos valores.
*/
package main

import (
	"fmt"
)

// Recebe dois parametros n1 e n2 e retorna um int8
func somar(n1, n2 int8) int8 {
	return n1 + n2
}

// Função com multiplos retornos são otimas pra quando lidamos com erros.
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println("Função")
	}
	f("texto")

	// Quando chamamos a função por ele ter dois retornos precisamos setar duas variaveis pra inferir elas.Podemos usar _ se não formos usar uma das variaveis.
	resultadosSoma, _ := calculos(10, 15)

	fmt.Println(resultadosSoma)
}
