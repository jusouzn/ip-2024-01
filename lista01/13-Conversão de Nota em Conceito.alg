algoritmo "13 Conversao de Nota em Conceito"
var
    nota: real
    conceito: caractere

inicio
    escreva("Informe a nota: ")
    leia(nota)
    
    se nota<6 entao
        conceito <- "D"
    senao se (nota>=6) e (nota<7.5)
        conceito <- "C"
    senao se (nota>=7.5) e (nota<9)
        conceito <- "B"
        senao se (nota>=9) e (nota<=10)
        conceito <- "A"
    fimse
    escreval("NOTA = ", nota:2:1, " CONCEITO = ", conceito)


fimalgoritmo