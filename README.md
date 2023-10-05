# CURSO DE GOLANG
👨‍⚖️GO É UMA LINGUAGEM DE PROGRAMAÇÃO.

[![GitHub Repo stars](https://img.shields.io/badge/VILHALVA-GITHUB-03A9F4?logo=github)](https://github.com/VILHALVA) 
[![GitHub Repo stars](https://img.shields.io/badge/VEJA%20OS-VIDEOS-03A9F4?logo=youtube)](https://www.youtube.com/@vilhalva100/search?query=Golang)

[![GitHub Repo stars](https://img.shields.io/badge/VEJA-DOCUMENTAÇÃO-03A9F4?logo=google)](https://go.dev/doc/) <br>

[![GitHub Repo stars](https://img.shields.io/badge/-PLAYLIST%20DO%20YOUTUBE-blueviolet)](https://youtube.com/playlist?list=PLHPgIIn9ls6-1l7h8RUClMKPHi4NoKeQF&si=dQATi45Dly6a-hf9)

<img src="https://images.crunchbase.com/image/upload/c_lpad,h_256,w_256,f_auto,q_auto:eco,dpr_1/altgiuwyavczlcflyh8c" align="center" width="280"> <br>

![](https://i.imgur.com/waxVImv.png)

# 👀VISÃO PANORÂMICA:
| PERGUNTA | RESPOSTA |
| :---: | :---: |
| DATA DE CRIAÇÃO | 2007 |
| NOME DO CRIADOR | Robert Griesemer, Rob Pike e Ken Thompson | 
| SIGNIFICADO DO NOME | O nome "Go" foi escolhido como uma referência direta à velocidade de compilação e execução da linguagem |
| É BASEADA NO | C, Pascal, Algol |
| EXTENÇÃO DO ARQUIVO | .go |
| É MAIS USADA | Desenvolvimento web e sistemas embarcados |

- Data de criação: A linguagem Go foi criada no ano de 2007, mas seu desenvolvimento foi anunciado publicamente em novembro de 2009, quando o projeto Go foi lançado pelo Google.

- Nome do criador: A linguagem Go foi desenvolvida por três engenheiros do Google: Robert Griesemer, Rob Pike e Ken Thompson.

- Significado do nome: O nome "Go" foi escolhido como uma referência direta à velocidade de compilação e execução da linguagem, indicando que a linguagem foi projetada para ser rápida e eficiente.

- Baseada em: Go não é diretamente baseada em nenhuma linguagem específica, mas possui influências de várias linguagens, incluindo C, Pascal, Algol e outras. Ela foi projetada para ser uma linguagem de programação moderna, concisa e eficiente.

- Extensão do arquivo: Os arquivos de código-fonte em Go têm a extensão ".go". Por exemplo, um programa Go pode ser salvo com o nome "meu_programa.go".

- Uso principal: Go é uma linguagem de programação que tem sido amplamente adotada na construção de sistemas de alto desempenho, servidores web, aplicativos de rede, serviços em nuvem e muito mais. Ela é conhecida por sua concorrência nativa, o que a torna especialmente adequada para tarefas que envolvem comunicação em rede e processamento paralelo. Além disso, Go também é usada em desenvolvimento de sistemas embarcados e outras aplicações de baixo nível devido à sua eficiência e simplicidade.

# 🤳SINTAXE DA LINGUAGEM:
## 0) FUNDAMENTOS:
Aqui está um exemplo de código em Go que utiliza operadores aritméticos, relacionais e lógicos com tipos primitivos:

```go
package main

import "fmt"

func main() {
    // Declaração de variáveis
    var a, b int
    a = 10
    b = 5

    // Operadores aritméticos
    soma := a + b
    subtracao := a - b
    multiplicacao := a * b
    divisao := a / b
    resto := a % b

    fmt.Println("Operadores aritméticos:")
    fmt.Println("Soma:", soma)
    fmt.Println("Subtração:", subtracao)
    fmt.Println("Multiplicação:", multiplicacao)
    fmt.Println("Divisão:", divisao)
    fmt.Println("Resto:", resto)

    // Operadores relacionais
    igual := a == b
    diferente := a != b
    maior := a > b
    menor := a < b
    maiorOuIgual := a >= b
    menorOuIgual := a <= b

    fmt.Println("\nOperadores relacionais:")
    fmt.Println("Igual:", igual)
    fmt.Println("Diferente:", diferente)
    fmt.Println("Maior:", maior)
    fmt.Println("Menor:", menor)
    fmt.Println("Maior ou igual:", maiorOuIgual)
    fmt.Println("Menor ou igual:", menorOuIgual)

    // Operadores lógicos
    verdadeiro := true
    falso := false

    eLogico := verdadeiro && falso
    ouLogico := verdadeiro || falso
    negacao := !verdadeiro

    fmt.Println("\nOperadores lógicos:")
    fmt.Println("E lógico (true && false):", eLogico)
    fmt.Println("OU lógico (true || false):", ouLogico)
    fmt.Println("Negação (!true):", negacao)
}
```

Este código declara variáveis `a` e `b`, em seguida, realiza operações aritméticas, relacionais e lógicas com essas variáveis e exibe os resultados na saída. Lembre-se de que os operadores aritméticos são usados para realizar operações matemáticas, os operadores relacionais comparam valores e retornam valores booleanos (verdadeiro ou falso) e os operadores lógicos permitem combinar expressões booleanas de diferentes maneiras.

## 1) VARIAVEIS SIMPLES:
```go
package main

import "fmt"

func main() {
    // Declaração de variáveis
    var nome string
    var idade int
    var altura float64
    var casado bool

    // Atribuição de valores às variáveis
    nome = "João"
    idade = 30
    altura = 1.75
    casado = false

    // Exibindo os valores das variáveis
    fmt.Println("Nome:", nome)
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
    fmt.Println("Casado:", casado)

    // Modificando valores das variáveis
    nome = "Maria"
    idade = 25
    altura = 1.68
    casado = true

    fmt.Println("\nValores modificados:")
    fmt.Println("Nome:", nome)
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
    fmt.Println("Casado:", casado)
}
```

Neste exemplo, temos quatro variáveis simples:

1. `nome` - Uma variável do tipo string que armazena um nome.
2. `idade` - Uma variável do tipo int que armazena uma idade.
3. `altura` - Uma variável do tipo float64 que armazena uma altura.
4. `casado` - Uma variável do tipo bool que armazena um estado civil.

As variáveis são declaradas e, em seguida, são atribuídos valores a elas. Posteriormente, os valores são modificados e os novos valores são exibidos na saída usando a função `fmt.Println`.

Este é um exemplo simples de como declarar, atribuir e utilizar variáveis em Go.

Aqui está um exemplo simples em que o usuário digita um valor e o programa exibe esse valor:

```go
package main

import (
    "fmt"
)

func main() {
    var valor int

    // Solicita ao usuário que insira um valor
    fmt.Print("Digite um valor: ")
    _, err := fmt.Scan(&valor)
    if err != nil {
        fmt.Println("Erro ao ler o valor:", err)
        return
    }

    // Exibe o valor inserido pelo usuário
    fmt.Println("O valor digitado foi:", valor)
}
```

Neste programa, usamos `fmt.Scan` para ler um valor inteiro digitado pelo usuário. Em seguida, exibimos o valor lido na saída padrão usando `fmt.Println`. Se ocorrer algum erro na leitura do valor, o programa tratará o erro e exibirá uma mensagem de erro. Caso contrário, ele exibirá o valor inserido pelo usuário.

## 2) ESTRUTURA CONDICIONAL:
### ESTRUTURA IF-ELSE:
A estrutura `if-else` em Go é usada para controlar o fluxo do programa com base em uma condição. Ela permite que você execute um bloco de código se uma condição for verdadeira (`if`), e outro bloco de código se a condição for falsa (`else`). Aqui está a sintaxe básica da estrutura `if-else` em Go:

```go
if condição {
    // Código a ser executado se a condição for verdadeira
} else {
    // Código a ser executado se a condição for falsa
}
```

Aqui estão alguns exemplos de como usar `if-else` em Go:

```go
package main

import "fmt"

func main() {
    numero := -5

    if numero >= 0 {
        fmt.Println("O número é positivo.")
    } else {
        fmt.Println("O número é negativo.")
    }
}
```

Neste exemplo, verificamos se o valor da variável `numero` é maior ou igual a zero. Se for verdadeiro, exibimos "O número é positivo." Se for falso, exibimos "O número é negativo."
```go
package main

import "fmt"

func main() {
    numero := 7

    if numero%2 == 0 {
        fmt.Println("O número é par.")
    } else {
        fmt.Println("O número é ímpar.")
    }
}
```

Neste exemplo, usamos o operador `%` para verificar se o número é divisível por 2. Se o resultado for igual a zero, o número é par; caso contrário, é ímpar.

A estrutura `if-else` é fundamental para criar lógica condicional em programas Go e é usada para tomar decisões com base nas condições fornecidas.

### ESTRUTURA IF-ELSE, ELSE IF:
```go
package main

import "fmt"

func main() {
    idade := 20

    if idade < 18 {
        fmt.Println("Você é menor de idade.")
    } else if idade >= 18 && idade < 60 {
        fmt.Println("Você é adulto.")
    } else {
        fmt.Println("Você é idoso.")
    }
}
```

Neste exemplo, usamos `if-else if-else` para verificar várias condições em sequência. O primeiro bloco `if` verifica se a idade é menor que 18, o segundo bloco `else if` verifica se a idade está entre 18 e 59, e o último bloco `else` é executado se nenhuma das condições anteriores for verdadeira.

### ESTRUTURA SWITCH:
A estrutura `switch` em Go é usada para criar ramificações condicionais com base em um valor de expressão. É uma alternativa mais limpa e eficiente para uma série de instruções `if-else if-else`. Aqui está a sintaxe básica da estrutura `switch` em Go:

```go
switch expressao {
case valor1:
    // Código a ser executado se a expressão for igual a valor1
case valor2:
    // Código a ser executado se a expressão for igual a valor2
default:
    // Código a ser executado se nenhum dos casos anteriores for verdadeiro
}
```

Aqui estão alguns exemplos de como usar `switch` em Go:

```go
package main

import "fmt"

func main() {
    dia := "quarta-feira"

    switch dia {
    case "segunda-feira":
        fmt.Println("É segunda-feira.")
    case "terça-feira":
        fmt.Println("É terça-feira.")
    case "quarta-feira":
        fmt.Println("É quarta-feira.")
    case "quinta-feira":
        fmt.Println("É quinta-feira.")
    case "sexta-feira":
        fmt.Println("É sexta-feira.")
    default:
        fmt.Println("É fim de semana.")
    }
}
```

Neste exemplo, usamos `switch` para verificar o valor da variável `dia` e executar o código correspondente ao dia da semana.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    valor := 42

    switch reflect.TypeOf(valor).Kind() {
    case reflect.Int:
        fmt.Println("O valor é um inteiro.")
    case reflect.String:
        fmt.Println("O valor é uma string.")
    default:
        fmt.Println("O tipo do valor não é suportado.")
    }
}
```

Neste exemplo, usamos `switch` para verificar o tipo da variável `valor` usando a função `reflect.TypeOf`. Dependendo do tipo, executamos o código correspondente.

```go
package main

import "fmt"

func main() {
    numero := 2

    switch numero {
    case 1:
        fmt.Println("Um")
    case 2:
        fmt.Println("Dois")
        fallthrough // Permite que o controle flua para o próximo caso
    case 3:
        fmt.Println("Três")
    default:
        fmt.Println("Outro número")
    }
}
```

Neste exemplo, usamos `fallthrough` no caso 2, o que faz com que o controle flua para o próximo caso (caso 3). Portanto, a saída será "Dois" seguida de "Três".

A estrutura `switch` é uma ferramenta poderosa para controlar o fluxo do programa em Go e pode ser usada de forma eficaz para verificar valores e executar o código apropriado com base nas condições.

## 3) ESTRUTURA DE REPETIÇÃO:
### ESTRUTURA FOR:
A estrutura `for` em Go é usada para criar loops, permitindo que você execute um bloco de código repetidamente enquanto uma condição for verdadeira. Existem três formas principais de usar a estrutura `for` em Go: `for`, `for range` e `for...range` (também conhecido como loop infinito). Aqui estão os detalhes de cada uma delas:

```go
for inicialização; condição; pós-loop {
    // Código a ser repetido enquanto a condição for verdadeira
}
```

Exemplo:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

Neste exemplo, o loop `for` é usado para imprimir os números de 0 a 4.

```go
for {
    // Código a ser executado infinitamente
}
```

Exemplo:

```go
package main

import "fmt"

func main() {
    contador := 0
    for {
        fmt.Println("Executando infinitamente...")
        contador++
        if contador == 3 {
            break // Sai do loop após três iterações
        }
    }
}
```

Neste exemplo, usamos um loop `for` infinito e saímos dele manualmente com a instrução `break` após três iterações.

O `for range` é usado para iterar sobre elementos de uma coleção, como um array, slice, mapa, etc.

```go
for índice, valor := range coleção {
    // Código a ser executado para cada elemento da coleção
}
```

Exemplo com slice:

```go
package main

import "fmt"

func main() {
    numeros := []int{1, 2, 3, 4, 5}

    for índice, valor := range numeros {
        fmt.Printf("Índice: %d, Valor: %d\n", índice, valor)
    }
}
```

Neste exemplo, usamos `for range` para iterar sobre os elementos de um slice e exibir o índice e o valor de cada elemento.

A estrutura `for` é uma parte fundamental da programação em Go e permite que você controle a repetição de código de várias maneiras, tornando-a uma ferramenta versátil para lidar com iterações em seus programas.

### ESTRUTURA WHILE:
Em Go, não existe uma estrutura `while` explicitamente como em algumas outras linguagens de programação, como C ou Java. No entanto, você pode criar um loop com comportamento semelhante ao `while` usando a estrutura `for` com uma única condição. O resultado é uma construção que funciona como um loop `while`. Veja como você pode fazer isso:

```go
for condição {
    // Código a ser repetido enquanto a condição for verdadeira
}
```

Aqui está um exemplo de como você pode usar essa construção para criar um loop `while` em Go:

```go
package main

import "fmt"

func main() {
    contador := 0

    for contador < 5 {
        fmt.Println("Contador:", contador)
        contador++
    }
}
```

Neste exemplo, o loop será repetido enquanto a condição `contador < 5` for verdadeira. A cada iteração, o valor do contador é incrementado. Essa construção é equivalente a um loop `while` em outras linguagens de programação.

Lembre-se de que em Go, a estrutura `for` é bastante flexível e pode ser usada para criar diversos tipos de loops, incluindo loops `for` tradicionais, loops `while` e loops infinitos, dependendo de como você configura as condições e os comportamentos do loop.

## 4) VARIAVEIS COMPOSTAS:
### ARRAYS:
Em Go, um array é uma coleção fixa de elementos do mesmo tipo. Os arrays têm tamanho fixo e não podem ser alterados depois de criados. Aqui está como você pode declarar, inicializar e usar arrays em Go:

1. **Declaração de Arrays:**
   Para declarar um array em Go, você especifica o tipo dos elementos e o tamanho do array. A sintaxe básica é a seguinte:

   ```go
   var nomeDoArray [tamanho]TipoDosElementos
   ```

   Exemplo:

   ```go
   var numeros [5]int // Um array de inteiros com tamanho 5
   ```

2. **Inicialização de Arrays:**
   Você pode inicializar um array ao declará-lo, especificando os valores dos elementos entre chaves `{}`. O tamanho do array é deduzido automaticamente pelo número de elementos fornecidos.

   Exemplo:

   ```go
   numeros := [5]int{1, 2, 3, 4, 5}
   ```

   Você também pode deixar que Go determine o tamanho do array com base nos elementos fornecidos:

   ```go
   numeros := [...]int{1, 2, 3, 4, 5} // Go calcula o tamanho automaticamente
   ```

3. **Acessando Elementos de Arrays:**
   Os elementos de um array são acessados usando índices, onde o índice começa em 0 para o primeiro elemento e vai até `tamanho - 1` para o último elemento.

   Exemplo:

   ```go
   primeiroNumero := numeros[0] // Acessa o primeiro elemento
   segundoNumero := numeros[1]  // Acessa o segundo elemento
   ```

4. **Modificando Elementos de Arrays:**
   Você pode modificar os elementos de um array atribuindo novos valores a eles.

   Exemplo:

   ```go
   numeros[2] = 10 // Modifica o terceiro elemento para o valor 10
   ```

5. **Comprimento de Arrays:**
   Você pode obter o comprimento (número de elementos) de um array usando a função `len()`.

   Exemplo:

   ```go
   tamanho := len(numeros) // Retorna o comprimento do array
   ```

6. **Iterando por Arrays:**
   Você pode usar um loop `for` para iterar pelos elementos de um array.

   Exemplo:

   ```go
   for i := 0; i < len(numeros); i++ {
       fmt.Println(numeros[i])
   }
   ```

   Também é possível utilizar o `range` para percorrer arrays, que é uma abordagem mais simples e idiomaticamente Go:

   ```go
   for indice, valor := range numeros {
       fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
   }
   ```

Os arrays são úteis quando você precisa de uma coleção fixa de elementos, mas em Go, é mais comum usar slices, que são estruturas de dados mais flexíveis e dinâmicas que são construídas sobre arrays subjacentes. Slices permitem adicionar ou remover elementos, enquanto arrays têm tamanho fixo.

## DICIONÁRIOS:
Go não possui uma estrutura de dados chamada "dicionário" no sentido tradicional, como em algumas outras linguagens de programação, como Python, que oferece dicionários ou mapas. Em vez disso, Go possui uma estrutura de dados chamada "mapa" (ou "map") que é usada para mapear chaves a valores. Os mapas são a implementação de estrutura de dados chave-valor em Go.

Aqui está como você pode usar um mapa em Go:

1. **Declaração de Mapa:**
   Para declarar um mapa em Go, você usa a seguinte sintaxe:

   ```go
   var nomeDoMapa map[TipoDaChave]TipoDoValor
   ```

   Exemplo:

   ```go
   var notas map[string]int // Um mapa que mapeia nomes para notas (strings para inteiros)
   ```

2. **Inicialização de Mapa:**
   Os mapas precisam ser inicializados antes de serem usados. Você pode inicializá-los usando o operador `make` ou declarando e inicializando-os em uma única linha.

   Exemplo com `make`:

   ```go
   notas := make(map[string]int)
   ```

   Exemplo com declaração e inicialização:

   ```go
   notas := map[string]int{
       "Alice": 90,
       "Bob":   85,
       "Carol": 88,
   }
   ```

3. **Adicionando e Acessando Elementos:**
   Você pode adicionar elementos a um mapa usando a sintaxe `mapa[chave] = valor`. Para acessar um valor, você usa a chave correspondente.

   Exemplo:

   ```go
   notas["David"] = 92
   aliceNota := notas["Alice"]
   ```

4. **Verificando se uma Chave Existe:**
   Você pode verificar se uma chave existe em um mapa usando a seguinte sintaxe:

   ```go
   valor, existe := mapa[chave]
   ```

   Exemplo:

   ```go
   nota, existe := notas["Eva"]
   if existe {
       fmt.Printf("Nota de Eva: %d\n", nota)
   } else {
       fmt.Println("Eva não encontrada.")
   }
   ```

5. **Removendo Elementos:**
   Você pode remover um elemento de um mapa usando a função `delete`.

   Exemplo:

   ```go
   delete(notas, "Bob")
   ```

Os mapas em Go são uma maneira eficaz de associar chaves a valores e são amplamente usados em programação Go para implementar estruturas de dados semelhantes a dicionários ou dicionários propriamente ditos. Eles são úteis quando você precisa armazenar e recuperar informações com base em chaves exclusivas.

## 5) FUNÇÕES:
Em Go, funções são blocos de código reutilizáveis que realizam uma tarefa específica. As funções desempenham um papel fundamental na estrutura de um programa Go e são usadas para dividir o código em partes menores e mais gerenciáveis. Aqui está uma visão geral de como as funções funcionam em Go:

* **1.Declaração de Função:**

Para declarar uma função em Go, você usa a seguinte sintaxe:

```go
func nomeDaFuncao(parametro1 Tipo1, parametro2 Tipo2) TipoDeRetorno {
    // Corpo da função
    // Código para realizar uma tarefa
    return valorDeRetorno
}
```

- `nomeDaFuncao` é o nome da função.
- `parametro1`, `parametro2`, etc., são os parâmetros de entrada da função, cada um com seu tipo.
- `TipoDeRetorno` é o tipo de dado que a função retorna.
- `valorDeRetorno` é o valor que a função retorna usando a instrução `return`. A função pode retornar múltiplos valores.

* **2). Exemplo de Função Simples:**

Aqui está um exemplo simples de função que soma dois números inteiros:

```go
func soma(a int, b int) int {
    resultado := a + b
    return resultado
}
```

* **3) Chamada de Função:**

Para chamar uma função em Go, você simplesmente usa seu nome seguido dos argumentos entre parênteses. Se a função retorna um valor, você pode atribuí-lo a uma variável.

Exemplo:

```go
resultado := soma(5, 3)
fmt.Println(resultado) // Saída: 8
```

* **4) Funções Variádicas:**

Em Go, você pode criar funções variádicas, que podem aceitar um número variável de argumentos do mesmo tipo. Isso é feito usando `...` antes do tipo do parâmetro.

Exemplo:

```go
func somaVariadica(nums ...int) int {
    resultado := 0
    for _, num := range nums {
        resultado += num
    }
    return resultado
}
```

Você pode chamar essa função com um número qualquer de argumentos inteiros:

```go
resultado := somaVariadica(1, 2, 3, 4, 5)
fmt.Println(resultado) // Saída: 15
```

* **5) Funções Anônimas:**

Go suporta funções anônimas (também conhecidas como closures), que são funções sem nome que podem ser atribuídas a variáveis e chamadas posteriormente.

Exemplo:

```go
quadrado := func(x int) int {
    return x * x
}

resultado := quadrado(5)
fmt.Println(resultado) // Saída: 25
```

* **6) Passagem de Parâmetros por Valor e por Referência:**

Em Go, os parâmetros de função são passados por valor por padrão. Isso significa que uma cópia dos valores é passada para a função. No entanto, você pode passar um ponteiro para um valor para permitir que a função modifique o valor original.

* **7) Funções Recursivas:**

Go suporta funções recursivas, que são funções que chamam a si mesmas. No entanto, é importante usar a recursão com cuidado para evitar estouro de pilha.

* **8) Ponto de Entrada (Função `main`):**

Todo programa Go deve ter uma função `main`, que é o ponto de entrada do programa. O código dentro da função `main` é executado quando o programa é iniciado.

Exemplo:

```go
func main() {
    // Código principal do programa
}
```

As funções são uma parte fundamental da programação em Go e são usadas para organizar e reutilizar código de maneira eficaz. Elas desempenham um papel central na estrutura de um programa Go.

## 6) CLASS POO:
Em Go, não existe uma construção nativa de "classes" como em linguagens de programação orientada a objetos (OOP) tradicionais, como Java ou Python. Em vez disso, Go utiliza um modelo de programação baseado em structs e interfaces para alcançar os conceitos OOP. Vou mostrar um exemplo de como você pode aplicar os quatro pilares da programação orientada a objetos em Go usando structs e interfaces.

Aqui está um exemplo que demonstra encapsulamento, herança, polimorfismo e abstração:

```go
package main

import (
    "fmt"
    "math"
)

// Encapsulamento: campos privados em uma struct
type Circulo struct {
    raio float64
}

// Abstração: um método associado à struct
func (c Circulo) Area() float64 {
    return math.Pi * c.raio * c.raio
}

// Herança: uma struct embutida (Composição)
type Cilindro struct {
    Circulo    // Herda campos e métodos de Circulo
    altura float64
}

// Polimorfismo: Cilindro redefine o método Area
func (c Cilindro) Area() float64 {
    return 2 * math.Pi * c.raio * c.raio + 2 * math.Pi * c.raio * c.altura
}

func main() {
    circulo := Circulo{raio: 5.0}
    cilindro := Cilindro{
        Circulo: circulo,
        altura:  10.0,
    }

    // Abstração e Polimorfismo: uso do método Area em ambas as structs
    fmt.Printf("Área do Círculo: %.2f\n", circulo.Area())
    fmt.Printf("Área do Cilindro: %.2f\n", cilindro.Area())
}
```

Neste exemplo:

1. **Encapsulamento:** O campo `raio` da struct `Circulo` é privado (inicia com letra minúscula) e não é acessível fora da struct.

2. **Abstração:** O método `Area` é usado para calcular a área de um círculo e é associado à struct `Circulo`.

3. **Herança (Composição):** A struct `Cilindro` embute a struct `Circulo`, o que permite herdar seus campos e métodos.

4. **Polimorfismo:** A struct `Cilindro` redefine o método `Area`, permitindo que ela calcule sua própria área de acordo com suas características específicas.

Este exemplo ilustra como os conceitos de OOP podem ser implementados em Go usando structs e interfaces para alcançar encapsulamento, herança, polimorfismo e abstração. É importante notar que Go adota um modelo de OOP mais leve e focado na composição de tipos, em vez de herança tradicional de classes.

# 💖CARACTERISTICAS DA LINGUAGEM:
## ❤POSITIVAS:
1. **Sintaxe Clara e Simples:** Go possui uma sintaxe concisa e fácil de aprender, o que facilita a leitura e escrita de código. A linguagem enfatiza a simplicidade e legibilidade.

2. **Eficiência:** Go foi projetada com foco na eficiência e desempenho. Ela possui uma coleta de lixo eficiente, suporte para concorrência nativa e baixo consumo de recursos.

3. **Concorrência e Paralelismo:** Go possui primitivas de concorrência incorporadas, como goroutines e canais, que facilitam a criação de programas concorrentes e paralelos. Isso é especialmente útil para sistemas que precisam lidar com alta concorrência.

4. **Compilação Rápida:** O compilador Go é notavelmente rápido, o que acelera o desenvolvimento e a implantação de aplicativos.

5. **Segurança:** Go promove práticas seguras de programação, como verificação de limites de arrays e prevenção de vazamentos de memória.

6. **Facilidade de Manutenção:** A simplicidade da linguagem, juntamente com o forte sistema de tipos, facilita a manutenção e o refatoramento de código.

7. **Facilidade de Leitura e Documentação:** A linguagem possui suporte integrado para documentação e geração de documentação de código, tornando mais fácil para os desenvolvedores escreverem e compartilharem documentação legível.

8. **Ecossistema Forte:** Go possui uma comunidade de desenvolvedores ativa e um ecossistema robusto de bibliotecas e frameworks. Há muitas bibliotecas de código aberto disponíveis para diversas tarefas.

9. **Suporte para Interfaces:** Go possui interfaces que permitem a implementação de polimorfismo, permitindo que as estruturas de dados sejam flexíveis e adaptáveis.

10. **Cross-Platform:** Go é uma linguagem multiplataforma que compila de maneira eficiente para várias arquiteturas e sistemas operacionais.

11. **Backed by Google:** Go é mantido pelo Google, o que aumenta a confiança na sua continuidade e suporte.

12. **Ótimo para Servidores e Microserviços:** Go é amplamente utilizado em servidores e para criar microserviços devido à sua eficiência e capacidade de lidar com grande carga de tráfego.

## 🖤NEGATIVAS:
1. **Falta de Genéricos:** Go ainda não tinha suporte nativo para tipos genéricos. Isso pode levar a repetição de código em algumas situações e tornar o código menos flexível.

2. **Goroutines Mal Gerenciadas:** Embora as goroutines sejam uma característica poderosa para concorrência em Go, o programador deve garantir que elas sejam corretamente gerenciadas para evitar vazamentos de goroutines ou panics devido a erros de programação.

3. **Ausência de Herança de Classes:** Go não suporta herança de classes no sentido tradicional da programação orientada a objetos. Em vez disso, ele usa composição de tipos. Isso pode parecer limitante para desenvolvedores que estão acostumados com linguagens que têm herança de classes.

4. **Curva de Aprendizado para Novos Desenvolvedores:** A sintaxe única e as convenções de Go podem ser desafiadoras para novos desenvolvedores acostumados com outras linguagens. Além disso, a falta de recursos como herança de classes pode requerer uma mudança na forma de pensar sobre o design do código.

5. **Falta de Frameworks Abundantes:** Em comparação com algumas outras linguagens, Go pode ter menos frameworks disponíveis para tarefas específicas, o que pode aumentar o esforço necessário para construir aplicativos completos.

6. **Compatibilidade com Versões Anteriores:** Go prioriza a compatibilidade com versões anteriores, o que pode levar a algumas limitações no desenvolvimento de novos recursos e atrasos na adoção de práticas mais modernas.

7. **Problemas de Dependência:** A gestão de dependências em Go, embora tenha melhorado ao longo do tempo com ferramentas como o módulo Go, pode ser desafiadora em alguns cenários complexos.

8. **Ausência de Frameworks GUI Nativos:** Go não oferece suporte nativo para o desenvolvimento de aplicativos com interfaces gráficas de usuário (GUI), o que limita sua aplicação em certos tipos de aplicativos de desktop.

