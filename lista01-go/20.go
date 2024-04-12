package main

import (
	"fmt"
)

var h, m, s, segundos int

func main() {
	fmt.Print("Digite quantas horas: ")
	fmt.Scan(&h)
	fmt.Print("Digite quantos minutos: ")
	fmt.Scan(&m)
	fmt.Print("Digite quantos segundos: ")
	fmt.Scan(&s)

	segundos = h*3600 + m*60 + s
	fmt.Printf(" o tempo em segundos é = %v", segundos)
}
