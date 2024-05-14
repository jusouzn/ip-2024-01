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
	for n <= 0 || n > 1000 {
		fmt.Scan(&n)
	}
	V := criarVetor(n)

	var k int
	fmt.Scan(&k)
	maiores := 0
	for i := 0; i < len(V); i++ {
		if V[i] >= k {
			maiores++
		}
	}

	fmt.Print(maiores)

}
