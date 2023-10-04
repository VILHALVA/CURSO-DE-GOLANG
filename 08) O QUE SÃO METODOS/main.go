package main

import "fmt"

// Definindo uma struct chamada "Pessoa"
type Pessoa struct {
    Nome  string
    Idade int
}

// Método associado à struct "Pessoa" para imprimir os detalhes da pessoa
func (p Pessoa) Apresentacao() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

// Método associado à struct "Pessoa" para aumentar a idade
func (p *Pessoa) Aniversario() {
    p.Idade++
}

func main() {
    // Criando uma instância da struct "Pessoa"
    pessoa := Pessoa{
        Nome:  "Alice",
        Idade: 25,
    }

    // Chamando o método "Apresentacao" da instância da struct "Pessoa"
    pessoa.Apresentacao()

    // Chamando o método "Aniversario" para aumentar a idade em 1 ano
    pessoa.Aniversario()

    // Chamando novamente o método "Apresentacao" após o aniversário
    pessoa.Apresentacao()
}
