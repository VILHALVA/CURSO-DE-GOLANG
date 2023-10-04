package main
import "fmt"

func ola(nome string) {
	fmt.Println("OLÁ ",nome)
}

func soma(n1, n2 int64) int64 {
	return n1 + n2
}

func main() {
	ola("VILHALVA")
	fmt.Println("A SOMA DOS VALORES É",soma(67,89))
}