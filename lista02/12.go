package main

import "fmt"

func main() {
	var ValorIngresso, ValorInicial, ValorFinal float64
	fmt.Scan(&ValorIngresso, &ValorInicial, &ValorFinal)
	if ValorInicial >= ValorFinal {
		fmt.Println("INTERVALO INVALIDO.")
	} else {
		Inicial := ValorInicial
		for i := ValorIngresso; i >= ValorInicial; i-- {
			Diferenca := ValorIngresso - Inicial
			Valor := Inicial
			NumIng := (Diferenca * 50) + 120
			Despesas := 200 + NumIng*0.05
			Lucro := Valor*NumIng - Despesas
			fmt.Printf("V: %0.2f, N: %0.0f, L: %0.2f\n", Valor, NumIng, Lucro)
			Inicial++
		}

		Final := Inicial
		for i := ValorIngresso; i < ValorFinal; i++ {
			Diferenca := Final - ValorIngresso
			Valor := Final
			NumIng := 120 - (Diferenca * 60)
			Despesas := 200 + NumIng*0.05
			Lucro := Valor*NumIng - Despesas
			fmt.Printf("V: %0.2f, N: %0.0f, L: %0.2f\n", Valor, NumIng, Lucro)
			Final++
		}
	}
}

//serão vendidos 120 ingressos e que as despesas fixas serão de R$ 200,00 mais R$ 0,05 por cada ingresso. Diminuindo-se R$ 0,50 o preço dos ingressos, esperase que as vendas aumentem em 25 ingressos. Aumentando-se R$ 0,50 o preço dos ingressos, espera-se que as vendas diminuam 30 ingressos
