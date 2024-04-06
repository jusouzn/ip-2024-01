algoritmo "18 Soma de progressão aritmética"
var
    inicial, r, n, soma: inteiro

inicio
    escreva("Digite o valor inicial: ")
    leia(inicial)
    escreva("Digite a razão: ")
    leia(r)
    escreva("Digite o número de elementos: ")
    leia(n)

    enquanto n>=0 faca
        soma <- soma+inicial
        inicial <- inicial+r
        n <- n-1 
    fimenquanto   
    escreval(soma)   
fimalgoritmo