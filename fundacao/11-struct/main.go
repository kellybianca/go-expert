package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Pessoa interface {
	Desativar()
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	kelly := Client{
		Nome:  "kelly",
		Idade: 22,
		Ativo: true,
	}
	kelly.Cidade = "Maceió"
	kelly.Desativar()
	fmt.Println(kelly.Nome)
	fmt.Println(kelly.Idade, kelly.Ativo)
	Desativacao(kelly)
}
