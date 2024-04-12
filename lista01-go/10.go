package main

import (
	"fmt"
)

var a, b, c, d, x float64

func main() {
	fmt.Print("Informe o número a: ")
	fmt.Scan(&a)
	fmt.Print("Informe o número b: ")
	fmt.Scan(&b)
	fmt.Print("Informe o número c: ")
	fmt.Scan(&c)
	fmt.Print("Informe o número d: ")
	fmt.Scan(&d)

	x = a*d - b*c

	fmt.Printf("O VALOR DO DETERMINANTE É = %0.2f", x)

}
