# O QUE É GOROUTINES?
Goroutines são uma característica fundamental da linguagem de programação Go (também conhecida como Golang) que permitem a execução concorrente de tarefas de forma eficiente e com baixo consumo de recursos. Uma goroutine é uma unidade leve de execução que é gerenciada pelo runtime de Go e é muito mais eficiente em termos de consumo de memória e recursos do que threads tradicionais em sistemas operacionais.

Aqui estão algumas características importantes das goroutines:

1. **Leveza:** As goroutines são leves em comparação com as threads do sistema operacional. Você pode criar muitas goroutines em um programa Go sem sobrecarregar o sistema.

2. **Independência:** Cada goroutine é independente e executa uma função ou parte do código separadamente. Elas não compartilham memória diretamente, o que ajuda a evitar problemas de concorrência.

3. **Comunicação:** Goroutines podem se comunicar entre si usando canais (channels), que são estruturas de dados que permitem a troca segura de informações entre goroutines.

4. **Concorrência:** As goroutines permitem que você aproveite a concorrência em seu programa, permitindo que tarefas sejam executadas de forma paralela. Isso pode melhorar o desempenho e a capacidade de resposta de aplicativos.

5. **Facilidade de uso:** Go oferece uma sintaxe simples para criar e gerenciar goroutines usando a palavra-chave `go`, o que torna a programação concorrente mais acessível em comparação com outras linguagens.

Aqui está um exemplo simples de como criar e usar uma goroutine:

```go
package main

import (
    "fmt"
    "time"
)

func imprimirNumeros() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Número:", i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    // Inicia uma goroutine para imprimir números
    go imprimirNumeros()

    // Continua a execução do programa principal
    fmt.Println("Programa principal continuando...")

    // Aguarda um tempo suficiente para que a goroutine termine
    time.Sleep(6 * time.Second)
}
```

Neste exemplo, criamos uma goroutine para imprimir números em paralelo com o programa principal. Isso permite que o programa principal continue a executar outras tarefas enquanto a goroutine executa em segundo plano. As goroutines são uma parte fundamental da programação concorrente em Go e são amplamente utilizadas para criar aplicativos eficientes e responsivos.