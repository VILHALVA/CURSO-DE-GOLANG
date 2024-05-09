package main

import (
    "fmt"
)

// Definição da estrutura Pessoa
type Pessoa struct {
    Nome  string
    Idade int
}

// Método para a estrutura Pessoa
func (p Pessoa) Apresentar() {
    fmt.Printf("Olá, eu sou %s e tenho %d anos.\n", p.Nome, p.Idade)
}

// Definição da estrutura Estudante, que herda de Pessoa
type Estudante struct {
    Pessoa         // Incorporação de Pessoa para alcançar a herança
    Curso  string  // Novo campo específico para Estudante
}

// Método para a estrutura Estudante
func (e Estudante) Estudar() {
    fmt.Printf("%s está estudando no curso de %s.\n", e.Nome, e.Curso)
}

// Interface para implementação de métodos comuns
type Apresentavel interface {
    Apresentar()
}

func ApresentarTudo(a ...Apresentavel) {
    for _, item := range a {
        item.Apresentar()
    }
}

func main() {
    // Entrada de dados para criar uma instância da estrutura Estudante
    var nome, curso string
    var idade int

    fmt.Println("Informe o nome do estudante:")
    fmt.Scanln(&nome)

    fmt.Println("Informe a idade do estudante:")
    fmt.Scanln(&idade)

    fmt.Println("Informe o curso do estudante:")
    fmt.Scanln(&curso)

    // Criando uma instância da estrutura Estudante com os dados fornecidos
    estudante1 := Estudante{
        Pessoa: Pessoa{
            Nome:  nome,
            Idade: idade,
        },
        Curso: curso,
    }

    // Chamando o método Apresentar() da instância estudante1
    estudante1.Apresentar()

    // Chamando o método Estudar() da instância estudante1
    estudante1.Estudar()

    // Usando polimorfismo para chamar Apresentar() em diferentes tipos de estruturas
    ApresentarTudo(estudante1, Pessoa{"José", 25})
}
