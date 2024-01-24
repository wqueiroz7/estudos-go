package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"pacotes/auxiliar"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliar.Escrever()

	checkmail.ValidateFormat("devbook@gmail.com")

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)

	// Go mod tidy, remove todas dependencias não usadas
}
