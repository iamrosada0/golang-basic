package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("Hello  %s %s\n", "world", "GO")

// 	fmt.Printf("2 + c5 = %v \n", 2+5)
// }

//CONSTANT
import "fmt"

const a = 10

func main() {
	//declarando multiplas variaves
	var d, m, e, f = 4, 5, 6, 7
	//variaveis, seus valores podem ser alterado
	//const nao se pode alterar o valor

	fmt.Printf("2 + %v", 2+a)
	// \n serve para quebrar a linha
	fmt.Printf("valor de d= %v\n, valor de m= %v \n,valor de e= %v \n, valor de f= %v \n", d, m, e, f)
}
