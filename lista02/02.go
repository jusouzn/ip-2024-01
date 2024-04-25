package main

import "fmt"

func main() {
	var popA, popB, anos int //popBsempre menor q popA

	const (
		taxA = 1.03
		taxB = 1.015
	)

	fmt.Scan(&popA)
	fmt.Scan(&popB)

	for i := 0; i >= 0; i++ {
		if popA == popB || popA > popB {
			anos = i
			break
		} else {
			popA = int(float64(popA) * taxA)
			popB = int(float64(popB) * taxB)
		}
	}

	fmt.Printf("ANOS = %d", anos)
}
