algoritmo "Consumo de Energia"
var
    salarioM, gastoKW, cadaKW, valorConsumo: real

inicio
    escreva("Digite o valor do salário mínimo: ")
    leia(salarioM)
    escreva("Digite a quantidade de kW gasta pela residência: ")
    leia(gastoKW)
    
    cadaKW <- (salarioM*0.7)/100
    valorConsumo <- cadaKW*gastoKW


    escreval("Custo por kW: R$ ", cadaKW)
    escreval("Custo do consumo: R$ ", valorConsumo)
    escreval("Custo com desconto: R$ ", (valorConsumo*0.9))
fimalgoritmo