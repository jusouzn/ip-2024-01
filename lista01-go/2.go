package main

import "fmt"

var CasoTeste int
var PublicoTotal, Popular, Geral, Arquibancada, Cadeiras, RendaTotal float64

func main() {
	fmt.Print("Digite o número de jogos: ")
	fmt.Scan(&CasoTeste)
	for i := 1; i <= CasoTeste; i++ {
		fmt.Print("Informe o número de pessoas: ")
		fmt.Scan(&PublicoTotal)
		fmt.Print("Informe a percentagem de Popular, Geral, Arquibancadas e Cadeiras respectivamente: ")
		fmt.Scan(&Popular, &Geral, &Arquibancada, &Cadeiras)
		RendaTotal = (PublicoTotal * (Popular / 100) * 1.00) + (PublicoTotal * (Geral / 100) * 5.00) + (PublicoTotal * (Arquibancada / 100) * 10.00) + (PublicoTotal * (Cadeiras / 100) * 20.00)
		fmt.Printf("A renda do jogo N. %d é = %.2f\n", i, RendaTotal)
	}
}
