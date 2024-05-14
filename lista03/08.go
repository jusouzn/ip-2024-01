package main

import (
	"fmt"
)

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	var binary string
	for ; n > 0; n = n / 2 {
		binary = fmt.Sprintf("%d", n%2) + binary
	}
	return binary
}

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		fmt.Println(toBinary(n))
	}
}
