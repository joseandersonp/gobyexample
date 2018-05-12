// Go tem suporte interno para o _retorno múltiplo de valores_.
// Esse recurso é usado frequentemente na idiomática do Go,
// por exemplo, para retornar ambos valores de resultado e
// erro de uma função.

package main

import "fmt"

// O (int, int) nesta assinatura de função mostra que a
// função retorna 2 `int`s.
func vals() (int, int) {
    return 3, 7
}

func main() {

    // Aqui usamos o retorno de dois valores diferentes
    // da função com _atribuição múltipla_.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // If you only want a subset of the returned values,
    // use the blank identifier `_`.

    // Se você quiser apenas um subconjunto dos valores
    // retornados, use o _identificador de ausência_ `_`.
    _, c := vals()
    fmt.Println(c)
}
