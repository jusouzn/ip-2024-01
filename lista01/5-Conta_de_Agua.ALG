algoritmo "Conta de Água"
var
    consumo, valor: real
    conta: inteiro
    consumidorT: caracter

inicio
    escreva("Digite a sua conta: ")
    leia(conta)
    escreva("Digite o seu consumo de água: ")
    leia(consumo)
    escreva("Qual sa categoria de consumidor?(C-COMERCIAL, I-INDUSTRIAL, R-RESIDENCIAL): ")
    leia(consumidorT)
    
    escolha consumidorT
        caso "R", "r":
            valor <- 5.00 + (0.05 * consumo)
        caso 'C', 'c':
            se consumo <= 80 entao
                valor <- 500.00
            senao
                valor <- 500.00 + (0.25 * (consumo - 80))
            fimse
        caso 'I', 'i':
            se consumo <= 100 entao
                valor <- 800.00
            senao
                valor <- 800.00 + (0.04 * (consumo - 100))
            fimse
        caso contrario:
            escreval("Tipo de consumidor inválido!")
    fimescolha

    escreval("Conta: ", conta)
    escreval("Valor da Conta: R$ ", valor)

fimalgoritmo