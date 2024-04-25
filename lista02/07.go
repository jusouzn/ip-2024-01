package main

import "fmt"

func main() {
	numeros := []int{}
	var mi, mp, ni, np float64

	for {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numeros = append(numeros, num)
	}

	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			mp += float64(numeros[i])
			np++
		} else {
			mi += float64(numeros[i])
			ni++
		}
	}
	fmt.Printf("MEDIA PAR: %0.2f\nMEDIA IMPAR: %0.2f\n", mp/np, mi/ni)
}
