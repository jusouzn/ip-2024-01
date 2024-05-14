package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	notas := make([]int, N)
	frequencia := make(map[int]int)
	var maiorNota, indiceMaiorNota int

	for i := 0; i < N; i++ {
		fmt.Scan(&notas[i])
		frequencia[notas[i]]++

		if i == 0 || notas[i] > maiorNota {
			maiorNota = notas[i]
			indiceMaiorNota = i
		}
	}

	ultimaNota := notas[N-1]
	fmt.Printf("Nota %d, %d vezes\n", ultimaNota, frequencia[ultimaNota])
	fmt.Printf("Nota %d, indice %d\n", maiorNota, indiceMaiorNota)
}
