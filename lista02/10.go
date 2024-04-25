package main

import "fmt"

type Funcionario struct {
	Matricula int
	SalarioH  float64
	Horas     float64
}

func main() {

	var funcinarios []Funcionario
	//esse for serve para a entrada de valores
	for {
		var matricula int
		fmt.Scan(&matricula)
		if matricula == 0 {
			break
		} else {
			var funcionario Funcionario
			funcionario.Matricula = matricula
			fmt.Scan(&funcionario.SalarioH, &funcionario.Horas)
			funcinarios = append(funcinarios, funcionario)
		}

	}
	for i := 0; i < len(funcinarios); i++ {

		SalarioF := (funcinarios[i].SalarioH) * (funcinarios[i].Horas)

		fmt.Printf("%d %0.2f\n", funcinarios[i].Matricula, SalarioF)
	}
}
