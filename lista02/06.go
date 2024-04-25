package main

import "fmt"

func main() {
	var n, sequencia, sMaior int
	fmt.Scanln(&n)

	numeros := make([]int, n)
	sequencia = 1
	sMaior = sequencia
	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	for i := 0; i < n-1; i++ {
		if numeros[i] == numeros[i+1] {
			sequencia++
			if sequencia >= sMaior {
				sMaior = sequencia
			}
		} else {
			sequencia = 1
		}
	}
	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos", sMaior)
}
