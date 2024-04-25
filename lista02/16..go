package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for h := 1; h <= n; h++ {
		for i := 1; i < n; i++ {
			for j := 1; j < n; j++ {
				if i*i+j*j == h*h && i <= j {
					fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, i, j)
				}
			}
		}
	}
}
