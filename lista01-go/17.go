package main

import (
	"fmt"
)

var x, y int

func main() {
	fmt.Print("Digite 1° número: ")
	fmt.Scan(&x)
	fmt.Print("Digite 2° número: ")
	fmt.Scan(&y)

	if x%2 != 0 {
		fmt.Println("O primeiro número não é par.")
	} else {
		for y > 0 {
			fmt.Print(x, " ")
			x += 2
			y -= 1
		}
	}
}
