package main

import "fmt"

// Definindo uma struct base chamada "Animal"
type Animal struct {
    Nome string
}

// Método da struct "Animal" para emitir um som genérico
func (a Animal) EmitirSom() {
    fmt.Println("O animal faz um som.")
}

// Definindo uma struct "Cachorro" que incorpora "Animal"
type Cachorro struct {
    Animal
    Raca string
}

// Método da struct "Cachorro" para emitir um latido
func (c Cachorro) Latir() {
    fmt.Println("Woof! Woof!")
}

func main() {
    // Criando uma instância de "Cachorro"
    meuCachorro := Cachorro{
        Animal: Animal{Nome: "Max"},
        Raca:   "Golden Retriever",
    }

    // Acessando campos e métodos de "Cachorro"
    fmt.Printf("Nome: %s, Raça: %s\n", meuCachorro.Nome, meuCachorro.Raca)
    meuCachorro.EmitirSom() // Chama o método da struct base "Animal"
    meuCachorro.Latir()     // Chama o método da struct "Cachorro"
}
