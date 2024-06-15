package main

import "fmt"

func expo(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * expo(x, n-1)
}

func main() {
	fmt.Println(expo(7, 2))
}
