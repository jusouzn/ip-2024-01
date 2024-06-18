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

func inverter(vet []int) []int {
	wat := make([]int, len(vet))
	j := 0
	for i := len(vet) - 1; i >= 0; i-- {
		wat[j] = vet[i]
		j++
	}
	return wat
}
func maiorElemento(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	maiorE := vet[0]
	for _, numero := range vet {
		if numero > maiorE {
			maiorE = numero
		}
	}
	return maiorE
}
func menorElemento(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	menorE := vet[0]
	for _, numero := range vet {
		if numero < menorE {
			menorE = numero
		}
	}
	return menorE
}

func main() {
	var N int
	fmt.Scan(&N)
	V := criarVetor(N)
	W := inverter(V)
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", V[i])
	}
	fmt.Println("")
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", W[i])
	}
	fmt.Println("")
	fmt.Println(maiorElemento(V))
	fmt.Println(menorElemento(W))

}
