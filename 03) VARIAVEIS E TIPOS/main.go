package main

import "fmt"

var (
	nome string
	idade int64
	dinheiro float64
	a, b, c int32
)
func main() {
	nome = "VILHALVA"
	idade = 27
	dinheiro = 455
	a = 22
	b = 45
	c = 12
	fmt.Println("OLÁ ", nome, "VOCÊ TEM ", idade, "VOCÊ TEM ", dinheiro, "EM DINHEIRO")
	fmt.Println(a, b, c)
}