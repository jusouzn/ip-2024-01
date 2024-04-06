algoritmo "1_Aprovado ou Reprovado"
var
    N1, N2, N3, quadrado, concatenacao: inteiro

inicio
    escreva("Digite 1° número: ")
    leia(N1)
    enquanto N1>10 faca
        escreva("Erro, digite novamente 1° número: ")
        leia(N1)
    fimenquanto

    escreva("Digite 2° número: ")
    leia(N2)
    enquanto N2>10 faca
        escreva("Erro, digite novamente 2° número: ")
        leia(N2)
    fimenquanto
    
    escreva("Digite 3° número: ")
    leia(N3)
    enquanto N3>10 faca
        escreva("Erro, digite novamente 3° número: ")
        leia(N3)
    fimenquanto
    
    concatenacao <- N1*100 + N2*10 + N3
    quadrado <- concatenacao**2
    escreval(concatenacao)
    escreval(quadrado)
fimalgoritmo