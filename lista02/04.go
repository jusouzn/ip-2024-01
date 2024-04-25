package main

import "fmt"

func main() {
	var n, in, K, s, incremento float64
	fmt.Scan(&n, &in, &K, &s)

	//soma
	println("Tabuada de soma:")
	incremento = in
	for i := 0.0; i < K; i++ {
		fmt.Printf("%0.2f + %0.2f = %0.2f\n", n, incremento, n+incremento)
		incremento = incremento + s
	}

	println("Tabuada de subtração:")
	incremento = in
	for i := 0.0; i < K; i++ {
		fmt.Printf("%0.2f - %0.2f = %0.2f\n", n, incremento, n-incremento)
		incremento = incremento + s
	}

	println("Tabuada de multiplicacao:")
	incremento = in
	for i := 0.0; i < K; i++ {
		fmt.Printf("%0.2f x %0.2f = %0.2f\n", n, incremento, n*incremento)
		incremento = incremento + s
	}

	println("Tabuada de divisao:")
	incremento = in
	for i := 0.0; i < K; i++ {
		fmt.Printf("%0.2f / %0.2f = %0.2f\n", n, incremento, n/incremento)
		incremento = incremento + s
	}

}
