package main

import (
	"fmt"
)

func somaDivisores(n int) int {
	soma := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if i*i != n {
				soma = soma + i + n/i
			} else {
				soma = soma + i
			}
		}
	}
	return soma
}

func numerosAmigos(n int) {
	cont := 0
	i := 1
	for cont < n {
		soma1 := somaDivisores(i)
		soma2 := somaDivisores(soma1)
		if soma2 == i && i < soma1 {
			fmt.Printf("(%d,%d)\n", i, soma1)
			cont++
		}
		i++
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	numerosAmigos(n)
}
