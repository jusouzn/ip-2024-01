algoritmo "17 Série de pares"
var
    x, y: inteiro

inicio
    escreva("Digite 1° Número: ")
    leia(x)
    escreva("Digite 2° Número: ")
    leia(y)

    se (x MOD 2 <> 0) entao
        escreval ("O PRIMEIRO NUMERO NAO E PAR.")
    senao 
        enquanto y>0 faca
            escreva(x, " ")
            x <- x+2
            y <- y-1 
        fimenquanto      
    fimse
fimalgoritmo