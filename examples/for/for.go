// `for` é a única declaração de loop do Go.
// Aqui estão três tipos básicos de loops `for`.

package main

import "fmt"

func main() {

    // A forma mais básica, com uma única condição.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // Uma clássico inicial/condição/pós loop `for`.
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // `for` sem uma condição executará um loop infinitamente
    // até encontrar um `break` ou `return` no escopo da função.
    for {
        fmt.Println("loop")
        break
    }

    // Você também pode usar `continue` para saltar para
    // próxima iteração do loop.
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
