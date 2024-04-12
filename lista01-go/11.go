package main

import (
	"fmt"
)

var numero int

func main() {
	fmt.Print("Informe o número: ")
	fmt.Scan(&numero)
	if numero%3 == 0 || numero%5 == 0 {
		fmt.Println("O número é divisível.")
	} else {
		fmt.Println("O número não é divisível.")
	}

}
