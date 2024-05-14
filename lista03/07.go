package main

import "fmt"

func criarVetor(n int) []int {
	V := make([]int, n)
	for i := 0; i < len(V); i++ {
		var num int
		fmt.Scan(&num)
		V[i] = num
	}
	return V
}

func buscaSequencial(V []int, elemento int) int {
	var Quantidade int
	for i := 0; i < len(V); i++ {
		if V[i] <= elemento {
			Quantidade++
		}
	}
	return Quantidade
}

func maiorValor(V []int) int {
	maior := 0
	for i := 0; i < len(V); i++ {
		if V[i] > maior {
			maior = V[i]
		}
	}
	return maior
}

func main() {

	Vetores := [][]int{}
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		} else {
			V := criarVetor(n)
			Vetores = append(Vetores, V)
		}
	}

	for i := 0; i < len(Vetores); i++ {
		maior := maiorValor(Vetores[i])
		for j := 0; j <= maior; j++ {

			Quant := buscaSequencial(Vetores[i], j)
			fmt.Printf("(%d) %d\n", j, Quant)
		}
	}
}
