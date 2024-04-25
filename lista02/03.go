package main

import "fmt"

func main() {
	var n, resultado int
	fmt.Scan(&n)
	resultado = n
	for i := 1; i < n; i++ {
		resultado = resultado * i
	}

	fmt.Printf("n! = %d", resultado)

}
