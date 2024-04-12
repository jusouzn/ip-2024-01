package main

import (
	"fmt"
)

var inicial, r, n, soma int

func main() {
	fmt.Print("Digite o valor inicial, razão e n° de elementos respectivamente: ")
	fmt.Scan(&inicial, &r, &n)
	for n > 0 {
		soma = soma + inicial
		inicial = inicial + r
		n--
	}
	fmt.Println(soma)
}
