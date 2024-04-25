package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("Campeonato Inválido!")
	}
	final := 1
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Printf("Final %d: Time%d X Time%d\n", final, i, j)
			final++
		}
	}
}
