package main

import (
	"fmt"
	"math"
)

func criarVetor(n int) []int {
	V := make([]int, n)
	for i := 0; i < len(V); i++ {
		var num int
		fmt.Scan(&num)
		V[i] = num
	}
	return V
}
func MaisProximo(V []int) (int, int) {
	maisP := 999999
	testes := 0
	for i := 0; i < len(V); i++ {
		for j := 1; j < len(V); j++ {
			if i >= j {
				continue
			}
			testes++
			dis := int(math.Abs(float64(V[i] - V[j])))
			if dis < maisP {
				maisP = dis
			}
		}

	}
	return maisP, testes

}

func main() {
	var T int
	fmt.Scan(&T)
	Vetores := make([][]int, T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		Vetores[i] = criarVetor(N)
	}

	for i := 0; i < T; i++ {
		Distancia, Testes := MaisProximo(Vetores[i])
		fmt.Printf("%d %d\n", Distancia, Testes)
	}
}
