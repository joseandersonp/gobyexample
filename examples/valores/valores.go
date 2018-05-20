// Go possui vários tipos de valores, incluindo strings,
// inteiros, pontos flutuantes, booleanos, etc. Aqui
// estão alguns exemplos básicos.

package main

import "fmt"

func main() {

    // Strings, que podem ser concatenadas com `+`.
    fmt.Println("go" + "lang")

    // Inteiros e pontos flutuantes.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Booleanos e com operadores booleanos.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
