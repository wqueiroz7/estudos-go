package auxiliar

import "fmt"

// A letra maiuscula e muito importante para que possamos importar esse funcao em outro pacote
// Se uma funcao começa com letra minuscula, significa que ela só é acessivel para dentro do pacote que ela ta.E maiuscula ela pode ser usada em outros pacotes

func Escrever() {
	fmt.Println("Pacote Auxiliar")
}
