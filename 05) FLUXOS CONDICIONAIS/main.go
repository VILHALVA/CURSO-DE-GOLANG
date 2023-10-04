package main

import "fmt"

func main() {
    // Suponha que temos duas notas
    nota1 := 85
    nota2 := 92

    // Calcula a média
    media := (nota1 + nota2) / 2

    // Verifica se a média é maior ou igual a 70
	fmt.Println("COM NOTA1:",nota1,"NOTA2:",nota2,"SUA MÉDIA É",media)
    if media >= 70 && media < 100{
        fmt.Println("Aprovado!")
    } else if media < 70 && media >= 50 {
        fmt.Println("Recuperação.")
    } else {
		fmt.Println("Reprovado!")
	}

    // Exemplo: 3 representa terça-feira
    diaDaSemana := 3 

    switch diaDaSemana {
    case 1:
        fmt.Println("Domingo")
    case 2:
        fmt.Println("Segunda-feira")
    case 3:
        fmt.Println("Terça-feira")
    case 4:
        fmt.Println("Quarta-feira")
    case 5:
        fmt.Println("Quinta-feira")
    case 6:
        fmt.Println("Sexta-feira")
    case 7:
        fmt.Println("Sábado")
    default:
        fmt.Println("Dia da semana inválido")
    }
}


