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

func BubbleSort(array []int) []int {
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
func main() {
	var N int
	fmt.Scan(&N)
	V := criarVetor(N)
	vetorOrganizado := BubbleSort(V)
	for i := 0; i < N; i++ {
		fmt.Printf("%d \n", vetorOrganizado[i])
	}
}
