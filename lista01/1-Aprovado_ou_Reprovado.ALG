algoritmo "1_Aprovado ou Reprovado"
var
    N1, N2, N3, MEDIA: real

inicio
    escreva("Digite nota 1: ")
    leia(N1)
    escreva("Digite nota 2: ")
    leia(N2)
    escreva("Digite nota 3: ")
    leia(N3)
    MEDIA <- (N1 + N2 + N3) / 3
    escreval("M�DIA: ", MEDIA)
    se (MEDIA >= 6) entao
        escreval("APROVADO")
    senao
        escreval("REPROVADO")
    fimse
fimalgoritmo