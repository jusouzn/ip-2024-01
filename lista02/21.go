package main

import "fmt"

func main() {
	var ordens []string
	for i := 0; i >= 0; i++ {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		} else {
			lista := make([]int, n)

			for i := 0; i < n; i++ {
				var k int
				fmt.Scan(&k)
				lista[i] = k
			}
			ok := true
			for i := 0; i < n-1; i++ {
				if lista[i] > lista[i+1] {
					ok = false
					break
				}
			}
			if ok {
				ordens = append(ordens, "ORGANIZADA")
			} else {
				ordens = append(ordens, "DESORGANIZADA")
			}
		}

	}
	for i := 0; i < len(ordens); i++ {
		fmt.Printf("%s \n", ordens[i])
	}
}
