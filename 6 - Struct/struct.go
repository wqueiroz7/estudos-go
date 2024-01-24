package main

/*
Structs são uma coleção de campos que cada campo tem um nome e um tipo, o Struct é um tipo
*/
type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {

	var usuario1 usuario
	usuario1.nome = "Wilson"
	usuario1.idade = 33

	/**/

}
