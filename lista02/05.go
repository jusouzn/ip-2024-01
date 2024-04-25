package main

import "fmt"

func main() {
	var n, sequencia, sMaior int
	fmt.Scanln(&n)

	numeros := make([]int, n)
	sequencia = 0
	sMaior = sequencia
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	for i := 0; i < n-1; i++ {
		if numeros[i] < numeros[i+1] {
			sequencia++
			if sequencia >= sMaior {
				sMaior = sequencia
			}
		} else {
			sequencia = 0
		}
	}
	fmt.Print("O comprimento do segmento crescente máximo é: ", sMaior)
}
