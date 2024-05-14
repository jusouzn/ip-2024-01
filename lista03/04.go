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

func main() {
	var n int
	fmt.Scan(&n)
	V := criarVetor(n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", V[i])
	}
}
