package main

import (
    "fmt"
    "sync"
)

func main() {
    numerosCh := make(chan int, 5)    // Usando um buffer de tamanho 5 para números
    letrasCh := make(chan string, 5) // Usando um buffer de tamanho 5 para letras

    var wg sync.WaitGroup
    done := make(chan struct{}) // Canal para sinalizar o término de todas as goroutines

    // Goroutine para enviar números
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(numerosCh)
        for i := 1; i <= 5; i++ {
            numerosCh <- i
        }
    }()

    // Goroutine para enviar letras
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(letrasCh)
        letras := "ABCDE"
        for _, letra := range letras {
            letrasCh <- string(letra)
        }
    }()

    // Goroutine para receber e imprimir números e letras
    wg.Add(1)
    go func() {
        defer wg.Done()
        for {
            select {
            case num, ok := <-numerosCh:
                if !ok {
                    fmt.Println("Channel de números fechado.")
                    close(done) // Sinaliza o término
                    return
                }
                fmt.Println("Número:", num)
            case letra, ok := <-letrasCh:
                if !ok {
                    fmt.Println("Channel de letras fechado.")
                    close(done) // Sinaliza o término
                    return
                }
                fmt.Println("Letra:", letra)
            }
        }
    }()

    // Goroutine para aguardar o término de todas as goroutines
    go func() {
        wg.Wait() // Aguarda todas as goroutines
        close(done) // Sinaliza o término
    }()

    // Aguarda o sinal de término
    <-done
}
