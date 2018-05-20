// [_Funções variádicas_](http://en.wikipedia.org/wiki/Variadic_function)
// podem ser chamadas com qualquer número de argumentos.
// Por exemplo, `fmt.Println` é uma função variádica comum.

package main

import "fmt"

// Aqui está uma função que obtém um número arbitrário
// de `int`s como argumentos.
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // Funções variádicas podem ser chamadas da da mesma
    // maneira com argumentos individuais.
    sum(1, 2)
    sum(1, 2, 3)

    // Se você já tiver slice com múltiplos valores, aplique-os
    // a uma função variádica usando `func(slice...)`.
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
