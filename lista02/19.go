package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	j := 1
	for i := 1; i <= n; i++ {
		fmt.Printf("%d*%d*%d = ", i, i, i)
		soma := 0
		for {
			soma += j
			fmt.Printf("%d", j)
			j += 2
			if soma != (i * i * i) {
				fmt.Printf("+")
			} else {
				break
			}
		}
		fmt.Printf("\n")
	}
}
