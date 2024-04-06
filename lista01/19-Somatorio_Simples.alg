algoritmo "19 Somatório Simples"
var
    n, i: inteiro
    somatorio: real

inicio
    escreva("Digite um número inteiro, positivo e maior que 1: ")
    leia(n)
    se n<=1 entao
        escreva("Numero invalido!, digite um número inteiro, positivo e maior que 1: ")
        leia(n)
    fimse
    i <- 1
    enquanto n>=1 faca
        somatorio <- somatorio + 1/i
        i <- i+1
        n <- n-1 
    fimenquanto   
    escreval(somatorio:10:6)   
fimalgoritmo