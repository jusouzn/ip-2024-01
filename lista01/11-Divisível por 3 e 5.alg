algoritmo "11 Divisível por 3 e 5"
var
    numero, delta : real

inicio
    escreva("Informe um número: ")
    leia(numero)
    se (numero MOD 3 = 0) ou (numero MOD 5 = 0) entao
        escreval("O NUMERO E DIVISIVEL")
    senao
        escreval("O NUMERO NAO E DIVISIVEL")
    fimse


fimalgoritmo