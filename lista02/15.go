package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)
	A := make([]int, t)
	B := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Scan(&A[i], &B[i])
	}
	for i := 0; i < t; i++ {
		if inverter(A[i]) >= inverter(B[i]) {
			fmt.Println(inverter(A[i]))
		} else {
			fmt.Println(inverter(B[i]))
		}
	}

}
func inverter(numero int) int {
	invertido := 0
	for numero > 0 {
		unidade := numero % 10
		invertido = invertido*10 + unidade
		numero = numero / 10
	}
	return invertido
}
