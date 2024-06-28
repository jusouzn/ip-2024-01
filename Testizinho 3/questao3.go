package main

import "fmt"

func bubbleSort(vetor []int) {
	n := len(vetor)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if vetor[j] > vetor[j+1] {
				vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	V := make([]int, n)
	for i := 0; i < len(V); i++ {
		var num int
		fmt.Scan(&num)
		V[i] = num
	}
	bubbleSort(V)
	fmt.Println(V)
}
