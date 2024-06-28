package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inverter(s string) string {
	delimitante := ""
	partes := strings.Split(s, delimitante)
	novastring := ""
	for i := len(s) - 1; i >= 0; i-- {
		novastring = novastring + partes[i]
	}
	return novastring
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)
	investida := inverter(frase)
	fmt.Println(investida)
}
