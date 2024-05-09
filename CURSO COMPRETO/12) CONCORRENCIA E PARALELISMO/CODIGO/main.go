package main

import (
    "fmt"
    "sync"
)

func main() {
    // Definindo o número de goroutines que serão executadas em paralelo
    numGoroutines := 5

    // Canal para receber os resultados das goroutines
    resultados := make(chan int, numGoroutines)

    // Usamos WaitGroup para sincronizar o término de todas as goroutines
    var wg sync.WaitGroup
    wg.Add(numGoroutines)

    // Lançando as goroutines
    for i := 0; i < numGoroutines; i++ {
        go func(i int) {
            defer wg.Done() // Marca a goroutine como completa quando termina
            // Simulando algum trabalho
            resultado := trabalho(i)
            // Enviando o resultado para o canal
            resultados <- resultado
        }(i)
    }

    // Função anônima para fechar o canal após o término de todas as goroutines
    go func() {
        wg.Wait()
        close(resultados)
    }()

    // Lendo os resultados das goroutines do canal
    soma := 0
    for resultado := range resultados {
        soma += resultado
    }

    // Exibindo o resultado final
    fmt.Println("Soma total:", soma)
}

// Função que simula algum trabalho
func trabalho(id int) int {
    // Aqui poderia ser feita alguma computação mais demorada
    return id * 2
}
