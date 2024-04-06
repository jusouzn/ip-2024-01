algoritmo "8 Custo da lata de cerveja"
var
    areacirculo, arealateral, raio, altura, custo : real

inicio
    escreva("Informe o raio em metros: ")
    leia(raio)
    escreva("Informe a altura em metros: ")
    leia(altura)
    areacirculo <-  3.14159*raio^2
    arealateral <- 2*3.14159*raio*altura
    custo <- (2*areacirculo + arealateral)*100

    escreval("O VALOR DO CUSTO E = ", custo:6:2)


fimalgoritmo