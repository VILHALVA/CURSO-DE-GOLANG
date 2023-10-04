# HERANÇA
Go não suporta herança de classe como algumas outras linguagens de programação orientada a objetos, como Java ou C++. Em vez disso, Go usa uma abordagem de composição, que permite criar tipos compostos (structs) incorporando outros tipos, em vez de herdar de classes.

No entanto, é possível criar um tipo composto que incorpora um ou mais campos de outro tipo, o que pode se assemelhar a herança. Vou mostrar um exemplo simplificado disso:

```go
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
```

Neste exemplo, temos uma struct base chamada `Animal`, que possui um campo `Nome` e um método `EmitirSom`. Em seguida, definimos uma struct `Cachorro` que incorpora `Animal` e tem um campo adicional `Raca` e um método `Latir`.

Ao criar uma instância de `Cachorro`, podemos acessar tanto os campos e métodos de `Cachorro` quanto os campos e métodos herdados de `Animal`. No entanto, vale ressaltar que não é uma herança tradicional de classes, mas sim uma composição de tipos. Go encoraja o uso de composição e interfaces em vez de herança de classes para criar abstrações e reutilização de código.