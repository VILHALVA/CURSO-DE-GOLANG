package main

import "fmt"

func main() {
    // Loop usando o "for" clássico
    for i := 0; i < 5; i++ {
        fmt.Println("Número:", i)
    }

    // Loop "while" usando "for" em Go
    contador := 0
    for contador < 5 {
        fmt.Println("Número:", contador)
        contador++
    }

    // Loop usando "range" para iterar por elementos de uma slice
    numeros := []int{1, 2, 3, 4, 5}
    for indice, valor := range numeros {
        fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
    }

    contador = 0 // Reinicializa o contador

    loop:
        fmt.Println("Número:", contador)
        contador++

        if contador < 5 {
            goto loop
        }
}
