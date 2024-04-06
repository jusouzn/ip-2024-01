algoritmo "20 Tempo em segundos"
var
    h, m, s, segundos: inteiro

inicio
    escreva("Digite quantas hora: ")
    leia(h)
    escreva("Digite quantos minutos: ")
    leia(m)
    escreva("Digite quantos segundos: ")
    leia(s)
    
    segundos <- h*3600 + m*60 +s

    escreval("O TEMPO EM SEGUNDOS E = ", segundos)   
fimalgoritmo