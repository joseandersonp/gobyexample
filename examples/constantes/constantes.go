// Go suporta _constantes_ de caracteres, strings,
// booleanos e valores numéricos.

package main

import "fmt"
import "math"

// `const` declara um valor constante.
const s string = "constante"

func main() {
    fmt.Println(s)

    // Uma declaração `const` pode aparecer no lugar de
    // qualquer declaração `var`.
    const n = 500000000

    // Expressões constantes realizam operações aritméticas
    // com precisão arbitrária.
    const d = 3e20 / n
    fmt.Println(d)

    // A constante numérica não possui nenhum tipo até que
    // seja atribuído a um, tal como por uma conversão explícita
    fmt.Println(int64(d))

    // Um número pode receber um tipo usando-o em um contexto que
    // requer um, como uma atribuição de variável ou chamada de
    // função. Por exemplo, aqui `math.Sin` espera um `float64`.
    fmt.Println(math.Sin(n))
}
