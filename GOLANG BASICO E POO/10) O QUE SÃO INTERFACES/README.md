# O QUE SÃO INTERFACES
Em Go (Golang), as interfaces são um conceito fundamental que permite definir um conjunto de métodos que um tipo deve implementar. As interfaces são usadas para criar contratos que permitem que diferentes tipos se comportem de maneira semelhante, mesmo que sejam completamente diferentes em termos de estrutura subjacente. Isso facilita a criação de código flexível e genérico.

Uma interface é definida como um conjunto de métodos sem a implementação real desses métodos. Qualquer tipo que implemente todos os métodos listados em uma interface é considerado como implementando essa interface. Aqui está a sintaxe básica para definir e usar uma interface em Go:

```go
// Definir uma interface com métodos
type MinhaInterface interface {
    Metodo1() tipoDeRetorno
    Metodo2(parametro tipoDeParametro) tipoDeRetorno
    // ...
}

// Definir um tipo que implementa a interface
type MeuTipo struct {
    // Campos do tipo
}

func (mt MeuTipo) Metodo1() tipoDeRetorno {
    // Implementação do Método1 para MeuTipo
}

func (mt MeuTipo) Metodo2(parametro tipoDeParametro) tipoDeRetorno {
    // Implementação do Método2 para MeuTipo
}
```

Aqui estão alguns pontos importantes sobre interfaces em Go:

1. **Implementação Implícita:** Em Go, não é necessário declarar explicitamente que um tipo implementa uma interface. Se um tipo possui todos os métodos exigidos por uma interface, ele é considerado como implementando essa interface automaticamente.

2. **Polimorfismo:** As interfaces em Go permitem que você use polimorfismo, ou seja, você pode tratar diferentes tipos que implementam a mesma interface de maneira uniforme.

3. **Múltiplas Interfaces:** Um tipo pode implementar várias interfaces. Isso permite que você combine funcionalidades de diferentes fontes.

4. **Interfaces Vazias:** As interfaces vazias, denotadas por `interface{}`, são usadas para representar qualquer tipo de valor. Elas são usadas em situações em que você deseja escrever código genérico.

Aqui está um exemplo de como usar uma interface em Go:

```go
package main

import "fmt"

// Definir uma interface
type Formatter interface {
    Format() string
}

// Definir tipos que implementam a interface
type Pessoa struct {
    Nome string
}

func (p Pessoa) Format() string {
    return p.Nome
}

type Numero int

func (n Numero) Format() string {
    return fmt.Sprintf("%d", n)
}

func Imprimir(f Formatter) {
    fmt.Println(f.Format())
}

func main() {
    pessoa := Pessoa{"Alice"}
    numero := Numero(42)

    Imprimir(pessoa) // Imprime "Alice"
    Imprimir(numero) // Imprime "42"
}
```

Neste exemplo, temos uma interface `Formatter` que exige que os tipos implementem o método `Format()`. As structs `Pessoa` e `Numero` implementam essa interface e, portanto, podem ser usadas como argumentos para a função `Imprimir`, que aceita qualquer valor que implemente a interface `Formatter`.

As interfaces são uma parte essencial da programação em Go e são amplamente usadas para criar código genérico e flexível. Elas são uma maneira poderosa de alcançar a abstração e o polimorfismo em Go.