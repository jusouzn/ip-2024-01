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

	var m int
	fmt.Scan(&m)
	C := criarVetor(m)

	R := make([]string, m)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(V); j++ {
			if C[i] == V[j] {
				R[i] = ("ACHEI")
			}
		}
		if R[i] != "ACHEI" {
			R[i] = ("NÃO ACHEI")
		}
	}

	for i := 0; i < len(C); i++ {
		fmt.Println(R[i])
	}

}
