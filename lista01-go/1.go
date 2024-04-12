package main

import "fmt"

var N1, N2, N3, MEDIA float32

func main() {
	fmt.Print("Digite a Nota 1, 2 e 3: ")
	fmt.Scan(&N1, &N2, &N3)
	MEDIA = (N1 + N2 + N3) / 3
	fmt.Printf("Média: %0.1f", MEDIA)
	fmt.Println("")
	if MEDIA >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
