package main

import (
	"fmt"
)

func fatoresPrimos(n int) []int {
	var fatores []int
	for n%2 == 0 {
		fatores = append(fatores, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			fatores = append(fatores, i)
			n = n / i
		}
	}
	if n > 2 {
		fatores = append(fatores, n)
	}
	return fatores
}

func main() {
	var numero int
	fmt.Print("Digite um número inteiro maior que 1: ")
	_, erro := fmt.Scan(&numero)
	for erro != nil || numero <= 1 {
		fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", numero)
		fmt.Print("Digite um número inteiro maior que 1: ")
		_, erro = fmt.Scan(&numero)
	}
	fatores := fatoresPrimos(numero)
	fmt.Printf("%d = ", numero)
	for i, fator := range fatores {
		if i != 0 {
			fmt.Print(" x ")
		}
		fmt.Print(fator)
	}
	fmt.Println()
}
