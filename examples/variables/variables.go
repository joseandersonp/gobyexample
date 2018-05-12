// Em Go, _variáveis_ são declaradas explicitamente e usadas
// pelo compilador como, por exemplo, checar o tipagem-correta
// das chamadas das funções.

package main

import "fmt"

func main() {

    // `var` declara 1 ou mais variáveis.
    var a = "initial"
    fmt.Println(a)

    // Você pode declarar muitas variáveis de uma só vez.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go irá inferir o tipo de variáveis inicializadas.
    var d = true
    fmt.Println(d)

    // Variáveis declaradas sem uma inicialização correspondente são
    // de _valor-zero_. Por exemplo, o valor zero para um `int` é` 0`.
    var e int
    fmt.Println(e)

    // A sintaxe `:=` é uma forma abreviada de declarar e inicializar uma
    // variável. Por exemplo, para `var f string = "short"` como neste caso.
    f := "short"
    fmt.Println(f)
}
