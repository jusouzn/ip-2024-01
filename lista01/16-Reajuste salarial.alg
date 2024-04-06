algoritmo "16 Reajuste Salarial"
var
    salario, reajuste: real

inicio
    escreva("Digite o salário: ")
    leia(salario)
    
    se (salario<=300) entao
        reajuste <- salario*1.5
    senao
        reajuste <- salario*1.3       
    fimse
    escreval("SALARIO COM REAJUSTE = ", reajuste:6:2)
fimalgoritmo