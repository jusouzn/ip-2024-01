algoritmo "12 Locadora de charretes"
var
    horas, valorTotal : real

inicio
    escreva("Informe o número de horas que a charrete foi localizada: ")
    leia(horas)

    se horas<3 entao
        valorTotal <- 5*horas
    senao
        valorTotal <- horas/3*10
    fimse
    escreval("O VALOR A PAGAR E = ", valorTotal:6:2)


fimalgoritmo