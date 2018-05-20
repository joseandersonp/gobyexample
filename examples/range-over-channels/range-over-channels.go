// Em um exemplo [anterior](range), vimos como `for` e
// `range` fornecem iteração sobre estruturas de dados
// básicas. Também podemos usar essa sintaxe para iterar
// os valores recebidos de um canal.

package main

import "fmt"

func main() {

    // Vamos iterar sobre dois valores no canal `queue`.
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // Este `range` itera sobre cada elemento conforme
    // é recebido de `queue`. Invocamos `close` para fechar
    // o canal acima, pois a iteração terminará depois de
    // receber os dois elementos.
    for elem := range queue {
        fmt.Println(elem)
    }
}
