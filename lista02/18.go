package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	i := 2
	k := 1
	for a != 1 || b != 1 || c != 1 {

		if a%i == 0 || b%i == 0 || c%i == 0 {
			fmt.Printf("%d %d %d :%d\n", a, b, c, i)
			k *= i
		}
		if a%i == 0 {
			a = a / i
		}
		if b%i == 0 {
			b = b / i
		}
		if c%i == 0 {
			c = c / i
		}
		if a%i != 0 && b%i != 0 && c%i != 0 {
			i++
		}

	}
	fmt.Printf("MMC: %d", k)
}
