# O QUE SÃO METODOS
Em Go (Golang), os métodos são funções que estão associadas a um tipo específico, chamado de "receiver type" (tipo receptor). Eles permitem que você associe comportamentos específicos a um tipo de dados definido pelo usuário, tornando o código mais orientado a objetos e modular.

Aqui está a sintaxe básica para definir um método em Go:

```go
func (receiver ReceiverType) NomeDoMetodo() {
    // Código do método
}
```

- `receiver` é uma variável que representa o valor no qual o método será chamado.
- `ReceiverType` é o tipo de dados ao qual o método está associado.
- `NomeDoMetodo` é o nome que você escolhe para o método.

Aqui está um exemplo simples de como definir e usar um método em Go:

```go
package main

import "fmt"

// Definir um tipo de dados chamado "Pessoa"
type Pessoa struct {
    Nome string
    Idade int
}

// Definir um método para o tipo "Pessoa"
func (p Pessoa) Apresentar() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
    // Criar uma instância da struct "Pessoa"
    pessoa1 := Pessoa{"Alice", 30}

    // Chamar o método "Apresentar" da instância
    pessoa1.Apresentar()
}
```

Neste exemplo, criamos um tipo de dados `Pessoa` com um campo `Nome` e `Idade`, e em seguida, definimos um método chamado `Apresentar` para o tipo `Pessoa`. O método `Apresentar` imprime uma saudação com o nome e a idade da pessoa.

Quando chamamos o método `Apresentar` na instância `pessoa1`, ele age sobre os campos dessa instância, como se fosse uma função normal. No entanto, ele tem acesso ao contexto e aos campos do receptor `Pessoa`.

Métodos em Go são uma maneira poderosa de encapsular lógica relacionada a tipos de dados específicos, tornando o código mais organizado e reutilizável. Eles são frequentemente usados para criar abstrações e fornecer interfaces de tipos, permitindo o desenvolvimento de programas mais flexíveis e legíveis.