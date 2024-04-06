algoritmo "9 Cálculo de Delta na equação de Báskara"
var
    A, B, C, delta : real

inicio
    escreva("Informe o coeficiete A: ")
    leia(A)
    escreva("Informe o coeficiete B: ")
    leia(B)
    escreva("Informe o coeficiete C: ")
    leia(C)

    delta <-  B^2 -4*A*C

    escreval("O VALOR DE DELTA E = ", delta:6:2)


fimalgoritmo