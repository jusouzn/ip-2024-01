package main

import (
	"fmt"
)

func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

func trocaOpostosSeMenor(vetor []int, n int) {
	for i := 0; i < n/2; i++ {
		oposto := n - 1 - i
		if vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}
func criarVetor(n int) []int {
	V := make([]int, n)
	for i := 0; i < len(V); i++ {
		var num int
		fmt.Scan(&num)
		V[i] = num
	}
	return V
}

func main() {
	var casosTeste int
	fmt.Scan(&casosTeste)

	Vetores := make([][]int, casosTeste)
	for i := 0; i < casosTeste; i++ {
		var n int
		fmt.Scan(&n)
		V := criarVetor(n)
		trocaOpostosSeMenor(V, n)
		Vetores[i] = V
	}

	for i := 0; i < casosTeste; i++ {
		fmt.Println(Vetores[i])
	}
}
