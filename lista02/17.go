package main

import "fmt"

type Mercadoria struct {
	CodM   int
	PrecoC float64
	PrecoV float64
	Quant  int
}

func main() {
	var mercadorias []Mercadoria
	for {
		var codigo int
		fmt.Scan(&codigo)

		if codigo == 0 {
			break
		} else {
			var mercadoria Mercadoria
			mercadoria.CodM = codigo
			fmt.Scan(&mercadoria.PrecoC, &mercadoria.PrecoV, &mercadoria.Quant)
			mercadorias = append(mercadorias, mercadoria)
		}
	}
	Lucro10 := 0
	Lucro20 := 0
	LucroMais := 0
	MaiorLucro := 0
	MaiorVenda := 0
	TotalC := 0.0
	TotalV := 0.0
	TotalLucro := 0.0
	TotalOriginal := 0.0

	for i := 0; i < len(mercadorias); i++ {

		Lucro := float64(mercadorias[i].Quant)*mercadorias[i].PrecoV - float64(mercadorias[i].Quant)*mercadorias[i].PrecoC

		if Lucro > float64(MaiorLucro) {
			MaiorLucro = mercadorias[i].CodM
		}
		if mercadorias[i].Quant > MaiorVenda {
			MaiorVenda = mercadorias[i].CodM
		}

		if Lucro < (float64(mercadorias[i].Quant)*mercadorias[i].PrecoC)*0.1 {
			Lucro10++
		} else if Lucro <= (float64(mercadorias[i].Quant)*mercadorias[i].PrecoC)*0.2 && Lucro > (float64(mercadorias[i].Quant)*mercadorias[i].PrecoC)*0.1 {
			Lucro20++
		} else {
			LucroMais++
		}

		TotalV = TotalV + float64(mercadorias[i].Quant)*mercadorias[i].PrecoV
		TotalC = TotalC + float64(mercadorias[i].Quant)*mercadorias[i].PrecoC
		TotalLucro += Lucro

	}
	PLucro := TotalLucro - TotalOriginal
	Porcentagem := (PLucro * 100) / TotalC
	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", Lucro10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% & menor ou igual a 20%%: %d\n", Lucro20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", LucroMais)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", MaiorLucro)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", MaiorVenda)
	fmt.Printf("Valor total de compras: %0.2f, valor total de vendas: %0.2f e percentual de lucro total: %0.2f %%\n", TotalC, TotalV, Porcentagem)
}
