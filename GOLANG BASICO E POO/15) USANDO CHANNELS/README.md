# USANDO CHANNELS
Channels (canais) são uma característica fundamental em Go (Golang) para a comunicação e sincronização entre goroutines, que são threads leves que executam concorrentemente em um programa Go. Os canais são usados para transmitir dados de uma goroutine para outra de forma segura e eficiente.

Aqui estão os conceitos básicos sobre como usar canais em Go:

### Criando um Canal

Você pode criar um canal usando a função `make`. A sintaxe básica é a seguinte:

```go
meuCanal := make(chan TipoDeDados)
```

Onde `TipoDeDados` é o tipo de dados que você pretende transmitir pelo canal.

### Enviando e Recebendo Dados

- Para enviar dados para um canal, você usa a seta esquerda `<-` com a sintaxe `canal <- valor`. Por exemplo:

    ```go
    meuCanal <- 42
    ```

- Para receber dados de um canal, você usa a seta direita `->` com a sintaxe `valor := <- canal`. Por exemplo:

    ```go
    dado := <-meuCanal
    ```

### Fechando um Canal

Você pode fechar um canal usando a função `close`. Isso é útil para indicar que não haverá mais dados enviados para o canal.

```go
close(meuCanal)
```

### Exemplo de Uso de Canais

Aqui está um exemplo simples de como usar canais para comunicar duas goroutines:

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    meuCanal := make(chan int)
    var wg sync.WaitGroup

    // Goroutine para enviar dados para o canal
    wg.Add(1)
    go func() {
        defer wg.Done()
        meuCanal <- 42
    }()

    // Goroutine para receber dados do canal
    wg.Add(1)
    go func() {
        defer wg.Done()
        dado := <-meuCanal
        fmt.Println("Recebido:", dado)
    }()

    // Aguardar as goroutines terminarem
    wg.Wait()

    // Fechar o canal quando não for mais necessário
    close(meuCanal)
}
```

Neste exemplo, criamos um canal `meuCanal` e duas goroutines. Uma goroutine envia o valor 42 para o canal, enquanto a outra recebe e imprime o valor do canal. Usamos a `sync.WaitGroup` para garantir que ambas as goroutines terminem antes de fechar o canal.

Os canais são uma ferramenta poderosa para lidar com a concorrência em Go e são usados para coordenar a execução de goroutines e transmitir dados entre elas de forma segura. Eles desempenham um papel fundamental na criação de programas concorrentes e paralelos em Go.