package main

import "fmt"

// Definindo uma interface chamada "Pessoa"
type Pessoa interface {
    Nome() string
}

// Definindo uma struct para Pessoa Física
type PessoaFisica struct {
    nome  string
    cpf   string
}

// Implementação do método "Nome" para Pessoa Física
func (pf PessoaFisica) Nome() string {
    return pf.nome
}

// Definindo uma struct para Pessoa Jurídica
type PessoaJuridica struct {
    nome        string
    cnpj        string
    razaoSocial string
}

// Implementação do método "Nome" para Pessoa Jurídica
func (pj PessoaJuridica) Nome() string {
    return pj.nome
}

func main() {
    // Criando instâncias de Pessoa Física e Pessoa Jurídica
    pessoa1 := PessoaFisica{nome: "João", cpf: "123.456.789-01"}
    pessoa2 := PessoaJuridica{nome: "Empresa XYZ", cnpj: "99.999.999/0001-99", razaoSocial: "XYZ Ltda."}

    // Criando uma slice de interfaces Pessoa
    pessoas := []Pessoa{pessoa1, pessoa2}

    // Iterando sobre as pessoas e obtendo seus nomes
    for _, pessoa := range pessoas {
        if pf, ok := pessoa.(PessoaFisica); ok {
            fmt.Printf("Nome da Pessoa Física: %s\n", pf.Nome())
        } else if pj, ok := pessoa.(PessoaJuridica); ok {
            fmt.Printf("Nome da Pessoa Jurídica: %s\n", pj.Nome())
        } else {
            fmt.Println("Tipo de Pessoa desconhecido")
        }
    }
}
