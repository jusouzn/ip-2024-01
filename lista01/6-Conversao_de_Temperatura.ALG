algoritmo "6 Conversão de Temperatura"
var
    quantidade, Contador : inteiro
    tempF, tempC : real

inicio
    escreva("Digite o número de temperaturas(Fahrenheit) que você quer calcular: ")
    leia(quantidade)
    Contador <- 1
    enquanto (Contador <= quantidade) faca
        escreva("Informe a temperatura em Fahrenheit: ")
        leia(tempF)

        tempC <- (5*(tempF-32))/9

        escreval(tempF," FAHRENHEIT EQUIVALE A ", tempC, " CELSIUS")
        Contador <- Contador + 1
    fimenquanto
fimalgoritmo