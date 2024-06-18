package main

import (
	"fmt"
)

func criarVetor(n int) []float64 {
	V := make([]float64, n)
	for i := 0; i < len(V); i++ {
		var num float64
		fmt.Scan(&num)
		V[i] = num
	}
	return V
}
func BubbleSort(array []float64) []float64 {
	for i := (len(array) - 1); i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				auxiliar := array[j+1]
				array[j+1] = array[j]
				array[j] = auxiliar
			}
		}
	}
	return array
}

func Mediana(V []float64) float64 {
	if len(V) == 0 {
		panic("Erro: slice V está vazio. Não é possível calcular a mediana.")
	}
	var Me float64
	if len(V)%2 == 1 {
		Me = V[len(V)/2]
	} else {
		m1 := V[(len(V)/2)-1]
		m2 := V[len(V)/2]

		Me = (m1 + m2) / 2
	}
	return Me
}

func main() {
	var N int
	fmt.Scan(&N)
	Vetor := criarVetor(N)
	Vetor = BubbleSort(Vetor)

	fmt.Printf("Mediana: %.2f \n", Mediana(Vetor))
}
