package main

import "fmt"

type Aluno struct {
	Matricula    int
	NotasProvas  [8]float64
	NotasListas  [5]float64
	NotaTrabalho float64
	Presenca     float64
}

//funcao para calcular medias das provas e listas
func calcularMedia(notas []float64) float64 {
	soma := 0.0
	for i := 0; i < len(notas); i++ {
		soma = soma + notas[i]
	}
	return soma / float64(len(notas))
}

func main() {
	const (
		numProvas      = 8
		numListas      = 5
		presencaMinima = 75
		cargaHoraria   = 128
	)

	var alunos []Aluno
	//esse for serve para a entrada de valores
	for {
		var matricula int
		fmt.Scan(&matricula)

		if matricula == -1 {
			break
		} else {

			var aluno Aluno
			aluno.Matricula = matricula

			for i := 0; i < numProvas; i++ {
				fmt.Scan(&aluno.NotasProvas[i])
			}
			for i := 0; i < numListas; i++ {
				fmt.Scan(&aluno.NotasListas[i])
			}

			fmt.Scan(&aluno.NotaTrabalho, &aluno.Presenca)
			alunos = append(alunos, aluno)
		}

	}

	//calcula as medias com a funcao define a situação e já faz os prints
	for i := 0; i < len(alunos); i++ {
		mediaProvas := calcularMedia(alunos[i].NotasProvas[:])
		mediaListas := calcularMedia(alunos[i].NotasListas[:])
		notaFinal := 0.7*mediaProvas + 0.15*mediaListas + 0.15*alunos[i].NotaTrabalho

		situacao := ""
		if notaFinal >= 6 && alunos[i].Presenca >= (presencaMinima*0.75) {
			situacao = "APROVADO"
		} else if notaFinal >= 6 && alunos[i].Presenca < (presencaMinima*0.75) {
			situacao = "REPROVADO POR FREQUENCIA"
		} else if notaFinal < 6 && alunos[i].Presenca >= (presencaMinima*0.75) {
			situacao = "REPROVADO POR NOTA"
		} else {
			situacao = "REPROVADO POR NOTA E POR FREQUENCIA"
		}

		fmt.Printf("Matricula: %d, Nota Final: %0.2f, Situação Final: %s\n", alunos[i].Matricula, notaFinal, situacao)
	}
}
