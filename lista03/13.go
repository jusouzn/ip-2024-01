package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	notas := make([]int, N)
	frequencia := make(map[int]int)

	for i := 0; i < N; i++ {
		fmt.Scan(&notas[i])
		frequencia[notas[i]]++
	}
	maiorFrequencia := 0
	menorNum := 999
	for _, nota := range notas {
		if maiorFrequencia < frequencia[nota] || (maiorFrequencia == frequencia[nota] && nota < menorNum) {
			maiorFrequencia = frequencia[nota]
			menorNum = nota
		}
	}
	fmt.Println(menorNum)
	fmt.Println(maiorFrequencia)

}
