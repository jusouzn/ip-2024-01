package main

import "fmt"

func main() {

	ind := []int{}
	maiores := []int{}

	for {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		} else {
			maior := -9999999
			indice := 0
			V := make([]int, n)

			for i := 0; i < len(V); i++ {
				var num int
				fmt.Scan(&num)
				V[i] = num

				if V[i] > maior {
					maior = V[i]
					indice = i
				}
			}
			maiores = append(maiores, maior)
			ind = append(ind, indice)
		}

	}

	for i := 0; i < len(maiores); i++ {
		fmt.Printf("%d %d\n", ind[i], maiores[i])
	}
}
