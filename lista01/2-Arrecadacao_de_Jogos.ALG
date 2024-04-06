algoritmo "2_ Arrecadação de Jogos"
var
    PublicoTotal, CasoTeste, Contador: inteiro
    Popular, Geral, Arquibancada, Cadeiras, RendaTotal: real

inicio
    escreva("Digite o número de jogos: ")
    leia(CasoTeste)
    Contador <- 1
    enquanto (Contador <= CasoTeste) faca
        escreva("Informe o número de pessoas: ")
        leia(PublicoTotal)
        escreva("Informe a percentagem (Popular): ")
        leia(Popular)
        escreva("Informe a percentagem (Geral): ")
        leia(Geral)
        escreva("Informe a percentagem (Arquibancada): ")
        leia(Arquibancada)
        escreva("Informe a percentagem (Cadeiras): ")
        leia(Cadeiras)

        RendaTotal <- (PublicoTotal * (Popular / 100) * 1.00) + (PublicoTotal * (Geral / 100) * 5.00) + (PublicoTotal * (Arquibancada / 100) * 10.00) + (PublicoTotal * (Cadeiras / 100) * 20.00)

        escreval("A RENDA DO JOGO N. ", Contador, " E = ", RendaTotal)
        Contador <- Contador + 1
    fimenquanto
fimalgoritmo