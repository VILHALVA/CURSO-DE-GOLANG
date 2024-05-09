# FLUXOS CONDICIONAIS
Os fluxos condicionais permitem que você controle o fluxo de execução de um programa com base em condições específicas. Em Go, você pode usar a estrutura `if`, `else if` e `else` para implementar fluxos condicionais. Aqui está uma visão geral:

1. **Estrutura `if`:**

   A estrutura `if` é usada para executar um bloco de código se uma determinada condição for verdadeira (avaliada como `true`).

   ```go
   if condicao {
       // Código a ser executado se a condição for verdadeira
   }
   ```

   Exemplo:

   ```go
   idade := 20
   if idade >= 18 {
       fmt.Println("Você é maior de idade")
   }
   ```

2. **Estrutura `if` com `else`:**

   Você pode usar `else` após um bloco `if` para executar um bloco de código quando a condição do `if` for falsa.

   ```go
   if condicao {
       // Código a ser executado se a condição for verdadeira
   } else {
       // Código a ser executado se a condição for falsa
   }
   ```

   Exemplo:

   ```go
   idade := 15
   if idade >= 18 {
       fmt.Println("Você é maior de idade")
   } else {
       fmt.Println("Você é menor de idade")
   }
   ```

3. **Estrutura `if` com `else if`:**

   Quando você tem várias condições para verificar, pode usar a estrutura `else if` após o `if` para testar várias condições em sequência.

   ```go
   if condicao1 {
       // Código a ser executado se a condição1 for verdadeira
   } else if condicao2 {
       // Código a ser executado se a condição2 for verdadeira
   } else {
       // Código a ser executado se nenhuma das condições for verdadeira
   }
   ```

   Exemplo:

   ```go
   nota := 85
   if nota >= 90 {
       fmt.Println("A nota é A")
   } else if nota >= 80 {
       fmt.Println("A nota é B")
   } else if nota >= 70 {
       fmt.Println("A nota é C")
   } else {
       fmt.Println("A nota é D")
   }
   ```

4. **Operador Ternário:**

   Go não possui um operador ternário como algumas outras linguagens, mas você pode usar uma expressão condicional para obter um comportamento semelhante.

   ```go
   resultado := valorCondicao ? valorSeVerdadeiro : valorSeFalso
   ```

   Exemplo:

   ```go
   idade := 20
   mensagem := ""
   if idade >= 18 {
       mensagem = "Maior de idade"
   } else {
       mensagem = "Menor de idade"
   }
   ```

Essas são as principais estruturas de controle de fluxo condicional em Go. Elas permitem que você execute diferentes partes do código com base em condições lógicas. Certifique-se de entender bem essas estruturas, pois elas são fundamentais para a programação em Go. 

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