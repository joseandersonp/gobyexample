// Go suporta [_funções anônimas_](https://en.wikipedia.org/wiki/Anonymous_function),
// que podem formar [_closures_](http://en.wikipedia.org/wiki/Closure_(computer_science)).
// Funções anônimas são úteis quando você deseja definir uma função sem precisar nomeá-la.

package main

import "fmt"

// Esta função `intSeq` retorna outra função, que definimos
// anonimamente no corpo de `intSeq`. A função retornada
// _envolve_ a variável `i` para formar um closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    // Chamamos `intSeq`, atribuindo o resultado (uma função)
    // à variável `nextInt`. Este valor do tipo função captura
    // seu próprio valor de `i`, que será atualizado toda vez
    // que chamarmos nextInt.
    nextInt := intSeq()

    // Veja o efeito do closure chamando `nextInt`
    // algumas vezes.
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // Para confirmar que o estado é exclusivo para essa
    // função específica, crie e teste um novo.
    newInts := intSeq()
    fmt.Println(newInts())
}
