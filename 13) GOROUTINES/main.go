package main

import (
    "fmt"
    "time"
)

func imprimirNumeros() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Número: %d\n", i)
        time.Sleep(1 * time.Second) // Aguarda 1 segundo
    }
}

func imprimirSegundos() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Segundo: %d\n", i)
        time.Sleep(1 * time.Second) // Aguarda 1 segundo
    }
}

func main() {
    // Inicia duas goroutines para imprimir números e segundos
    go imprimirNumeros()
    go imprimirSegundos()

    // Aguarda um tempo suficiente para que as goroutines terminem
    time.Sleep(6 * time.Second)
}
