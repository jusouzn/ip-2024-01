package main

import "fmt"

func main() {
	var n, soma int
	fmt.Scan(&n)

	fmt.Printf("%d = ", n)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			soma = soma + i
			fmt.Printf("%d ", i)
		} else {
			continue
		}

		if soma == n {
			fmt.Printf("= %d (NÚMERO PERFEITO)", soma)
			break
		} else if soma > n {
			fmt.Printf("= %d (NÚMERO IMPERFEITO)", soma)
			break
		} else {
			fmt.Printf("+ ")
		}
	}
}
