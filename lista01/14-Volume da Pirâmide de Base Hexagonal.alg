algoritmo "14 Volume da PirÂmide de Base Hexagonal"
var
    areaB, volume, altura, aresta: real

inicio
    escreva("Informe a altura: ")
    leia(altura)
    escreva("Informe a aresta: ")
    leia(aresta)

    areaB <- (3 * aresta^2 * 1.73205)/2
    volume <- (1/3) * areaB * altura

    escreval("O VOLUME DA PIRAMIDE E = ", volume:6:2, " METROS CUBICOS")


fimalgoritmo