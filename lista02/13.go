package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var milhos int
	milhos = (n*32 + n*64) - n
	fmt.Println(milhos)
}
