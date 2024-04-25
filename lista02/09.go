package main

import (
	"fmt"
	"math"
)

func verificar(numero int) bool {
	if numero < 2 {
		return false
	} else {
		for i := 2; i <= int(math.Sqrt(float64(numero))); i++ {
			if numero%i == 0 {
				return false
			}
		}
		return true
	}

}

func main() {
	var n int
	var primo bool
	fmt.Scanln(&n)

	primo = verificar(n)

	if primo == true {
		fmt.Println("PRIMO")
	} else {
		fmt.Println("NÃO PRIMO")
	}
}
