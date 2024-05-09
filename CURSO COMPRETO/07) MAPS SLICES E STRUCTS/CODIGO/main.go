package main

import (
    "fmt"
)

// Definindo uma struct para representar informações sobre estudantes
type Estudante struct {
    Nome   string
    Idade  int
    Nota   float64
    Matricula string
}

func main() {
    // Criando um mapa para armazenar informações de estudantes
    estudantes := make(map[string]Estudante)

    // Criando um slice para armazenar chaves do mapa (nomes dos estudantes)
    chaves := []string{}

    // Adicionando estudantes ao mapa
    estudante1 := Estudante{"Alice", 20, 85.5, "12345"}
    estudante2 := Estudante{"Bob", 22, 92.0, "67890"}
    estudantes["12345"] = estudante1
    estudantes["67890"] = estudante2

    // Adicionando chaves (matrículas) ao slice
    chaves = append(chaves, "12345")
    chaves = append(chaves, "67890")

    // Iterando sobre os estudantes usando o mapa
    fmt.Println("Informações dos Estudantes:")
    for matricula, estudante := range estudantes {
        fmt.Printf("Matrícula: %s, Nome: %s, Idade: %d, Nota: %.2f\n", matricula, estudante.Nome, estudante.Idade, estudante.Nota)
    }

    // Iterando sobre os estudantes usando o slice de chaves
    fmt.Println("\nChaves (Matrículas) dos Estudantes:")
    for _, matricula := range chaves {
        estudante, encontrado := estudantes[matricula]
        if encontrado {
            fmt.Printf("Matrícula: %s, Nome: %s, Idade: %d, Nota: %.2f\n", matricula, estudante.Nome, estudante.Idade, estudante.Nota)
        }
    }
}
