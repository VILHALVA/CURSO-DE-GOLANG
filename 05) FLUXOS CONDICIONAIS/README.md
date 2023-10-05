# AS CHAVES DE IF E ELSE DEVEM ESTÁ NA MESMA LINHA!
Em Go (Golang), a linguagem é estruturada de forma a exigir que as chaves do `if` e `else` estejam na mesma linha. Isso significa que você não pode colocar o `else` abaixo da chave de fechamento do `if`, como é comum em algumas outras linguagens de programação, como C/C++ ou JavaScript.

Em Go, a sintaxe correta para um bloco `if` e `else` seria a seguinte:

```go
if condição {
    // Código a ser executado se a condição for verdadeira
} else {
    // Código a ser executado se a condição for falsa
}
```

Por exemplo:

```go
if idade >= 18 {
    fmt.Println("Você é maior de idade.")
} else {
    fmt.Println("Você é menor de idade.")
}
```

Observe que as chaves do `if` e do `else` estão alinhadas na mesma linha. Isso é uma característica da sintaxe Go e ajuda a manter o código limpo e uniforme.

Portanto, em Go, não é permitido colocar o `else` abaixo da chave de fechamento do `if`, e a formatação correta é com as chaves alinhadas na mesma linha.