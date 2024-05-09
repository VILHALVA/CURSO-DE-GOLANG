package main

import "fmt"

// Definindo uma interface chamada "Animal" com um método "FazerSom"
type Animal interface {
    FazerSom() string
}

// Definindo uma struct chamada "Cachorro" que implementa a interface "Animal"
type Cachorro struct {
    Nome string
}

// Implementação do método "FazerSom" para "Cachorro"
func (c Cachorro) FazerSom() string {
    return "Woof! Woof!"
}

// Definindo uma struct chamada "Gato" que também implementa a interface "Animal"
type Gato struct {
    Nome string
}

// Implementação do método "FazerSom" para "Gato"
func (g Gato) FazerSom() string {
    return "Meow!"
}

// Função que recebe qualquer valor que implementa a interface "Animal" e faz o som
func FazerBarulho(animal Animal) {
    fmt.Println(animal.FazerSom())
}

func main() {
    cachorro := Cachorro{Nome: "Rex"}
    gato := Gato{Nome: "Whiskers"}

    // Chamando a função "FazerBarulho" com instâncias de "Cachorro" e "Gato"
    FazerBarulho(cachorro)
    FazerBarulho(gato)
}
