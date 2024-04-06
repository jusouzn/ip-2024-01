algoritmo "7 Conversões para o Sistema Métrico"
var
    tempF, tempC, polegadas, milimetros : real

inicio
    escreva("Informe a temperatura em Fahrenheit: ")
    leia(tempF)
    escreva("Informe a quantidade de chuva dada em polegadas: ")
    leia(polegadas)
    tempC <- (5*(tempF-32))/9
    milimetros <- polegadas*25.4

    escreval("O VALOR EM CELSIUS = ", tempC)
    escreval("A QUANTIDADE DE CHUVA E = " milimetros "mm")


fimalgoritmo