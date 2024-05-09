# FUNÇÕES
As funções são blocos de código reutilizável que realizam tarefas específicas. Em Go, as funções são fundamentais e são definidas da seguinte forma:

```go
func nomeDaFuncao(parametro1 Tipo1, parametro2 Tipo2) TipoDeRetorno {
    // Corpo da função
    // Código para realizar uma tarefa
    return valorDeRetorno
}
```

Aqui estão os elementos-chave:

1. `func`: É a palavra-chave usada para definir uma função em Go.
2. `nomeDaFuncao`: É o identificador que você escolhe para nomear sua função.
3. `(parametro1 Tipo1, parametro2 Tipo2)`: São os parâmetros da função. Você declara o nome e o tipo de cada parâmetro dentro dos parênteses. As funções podem ter zero, um ou vários parâmetros.
4. `TipoDeRetorno`: É o tipo de dado que a função retorna. Se a função não retornar nada, use `void` (ou `nil` em Go).
5. `return valorDeRetorno`: A função pode retornar um valor com a instrução `return`. O tipo do valor de retorno deve corresponder ao tipo especificado.

Aqui está um exemplo simples de uma função em Go:

```go
package main

import "fmt"

func somar(a int, b int) int {
    resultado := a + b
    return resultado
}

func main() {
    resultadoSoma := somar(5, 3)
    fmt.Println("5 + 3 =", resultadoSoma)
}
```

Neste exemplo, temos uma função chamada `somar` que recebe dois números inteiros como parâmetros, realiza a soma e retorna o resultado como um número inteiro. Na função `main`, chamamos a função `somar` com os valores 5 e 3 e imprimimos o resultado.

Além disso, em Go, é possível retornar múltiplos valores de uma função. Por exemplo:

```go
func dividirEresto(a int, b int) (int, int) {
    quociente := a / b
    resto := a % b
    return quociente, resto
}
```

Nesse caso, a função `dividirEresto` retorna dois valores inteiros: o quociente e o resto da divisão de `a` por `b`.

Lembre-se de que em Go, a convenção é usar nomes de funções e variáveis em formato camelCase (primeira palavra com letra minúscula e palavras subsequentes começando com letra maiúscula) e nomes descritivos para tornar seu código mais legível.

