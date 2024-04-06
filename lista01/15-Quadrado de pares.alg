algoritmo "15 Quadrado de pares"
var
    N, teste: real

inicio
    escreva("Digite um número entre 5 e 200: ")
    leia(N)
    enquanto (N<5) e (N>200) faca
        escreva("Erro, digite novamente um número entre 5 e 200: ")
        leia(N)
    fimenquanto
    teste <- 2
    enquanto teste<=N faca
        se (teste MOD 2 = 0) entao
            escreval(teste,"ˆ2 = ", teste^2)
        fimse
        teste <- teste+2
    fimenquanto
fimalgoritmo