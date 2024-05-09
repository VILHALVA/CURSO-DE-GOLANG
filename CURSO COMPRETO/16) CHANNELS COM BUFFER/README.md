# CHANNELS COM BUFFER
Em Go (Golang), você pode criar canais com buffer, que são canais capazes de armazenar um número limitado de valores antes de bloquear a gravação. Isso pode ser útil em situações em que você deseja controlar o fluxo de dados ou quando deseja evitar que uma goroutine bloqueie outras. Aqui está como criar e usar canais com buffer:

### Criando um Canal com Buffer

Você pode criar um canal com buffer usando a função `make` com um segundo argumento, que especifica o tamanho do buffer.

```go
meuCanal := make(chan TipoDeDados, tamanhoDoBuffer)
```

Onde `TipoDeDados` é o tipo de dados que você pretende transmitir e `tamanhoDoBuffer` é o tamanho do buffer do canal.

### Enviando Dados para um Canal com Buffer

Você pode enviar dados para um canal com buffer usando a seta esquerda `<-` da mesma maneira que faria com um canal sem buffer.

```go
meuCanal <- valor
```

No entanto, o envio só bloqueará quando o buffer estiver cheio. Até lá, você pode continuar enviando valores para o canal sem bloquear.

### Recebendo Dados de um Canal com Buffer

Você pode receber dados de um canal com buffer da mesma maneira que faria com um canal sem buffer.

```go
valor := <-meuCanal
```

O recebimento bloqueará até que haja um valor disponível no canal, a menos que o canal tenha sido fechado.

### Exemplo de Uso de um Canal com Buffer

Aqui está um exemplo simples de uso de um canal com buffer:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    meuCanal := make(chan int, 3) // Canal com buffer de tamanho 3

    go func() {
        for i := 0; i < 5; i++ {
            meuCanal <- i
            fmt.Println("Enviado:", i)
        }
        close(meuCanal)
    }()

    // Aguardar um curto período para a goroutine enviar alguns valores
    time.Sleep(2 * time.Second)

    // Receber os valores do canal
    for valor := range meuCanal {
        fmt.Println("Recebido:", valor)
    }
}
```

Neste exemplo, criamos um canal com buffer de tamanho 3 e uma goroutine que envia valores para o canal. Como o buffer tem tamanho 3, a goroutine pode enviar três valores sem bloquear. Em seguida, o programa principal começa a receber os valores do canal, que foram enviados pela goroutine.

Os canais com buffer são úteis em situações em que você deseja controlar a quantidade de dados que pode ser transmitida ou quando precisa evitar bloqueios em goroutines que enviam ou recebem dados. Eles são uma ferramenta poderosa para a concorrência em Go.