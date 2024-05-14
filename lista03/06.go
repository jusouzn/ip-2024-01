package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	V := make([]int, n)
	soma := 0
	for i := 0; i < len(V); i++ {
		var num int
		fmt.Scan(&num)
		V[i] = num
		soma += num
	}

	fmt.Printf("%d ", soma)

}
