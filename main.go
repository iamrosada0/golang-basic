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
	// declaring multiple variables
	var d, m, e, f = 4, 5, 6, 7
	// variables, their values can be changed
	// constants cannot be changed
	// another way to use var
	var (
		u = 2
		l = 1
	)
	fmt.Printf("u + l = %v + %v = %v\n", u, l, u+l)
	fmt.Printf("2 + %v = %v\n", 2+a, 2+a)
	// \n is used to break the line
	fmt.Printf("valor de d= %v\nvalor de m= %v\nvalor de e= %v\nvalor de f= %v\n", d, m, e, f)
}
