package main

import "fmt"

func main() {
	var n, e float64
	fmt.Scan(&n, &e)
	rk := 1.0 //inicial
	for {
		rk1 := (rk + n/rk) / 2 //formula
		erro := (n - rk1*rk1) * -1
		fmt.Printf("r: %0.9f, erro: %0.9f\n", rk1, erro)
		rk = rk1
		if erro <= e {
			break
		}
	}
}
