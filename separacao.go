package main

import "fmt"

// domain
// onde fica a nossa regra de negócio
// ele não conhece ninguém
type Pessoa struct {
	nome   string
	genero string
	idade  int
}

func NewPessoa(n, g string, i int) Pessoa {
	return Pessoa{nome: n, genero: g, idade: i}
}

//=============

// main
func main() {

	CriarPessoa()
}

// camada de aplicação
// que vai orquestrar a cominunicação com a camada de dominio
func criaPessoa(nome, genero string, idade int) Pessoa {
	return Pessoa{nome, genero, idade}
}

// camada de infra-estruta
// onde lhe será obrigado inserir os dados
// mas ele não conhece o dominio
func CriarPessoa() {
	nome := "Novais"
	idade := 12
	genero := "M"

	p := criaPessoa(nome, genero, idade)

	fmt.Println(p)
}
