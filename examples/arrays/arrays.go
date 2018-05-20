// Em Go, um _array_ é uma sequência enumerada de elementos
// de um tamanho específico.

package main

import "fmt"

func main() {

    // Aqui nós criamos um array `a` que irá conter exatamente
    // 5 `int`s. O tipo dos elementos e o tamanho são partes
    // do tipo de array. Por padrão, um array é de valor-zero,
    // que para `int`s significa `0`s.
    var a [5]int
    fmt.Println("emp:", a)

    // Podemos definir um valor em um índice usando a sintaxe
    // `array[índice] = valor` e obter um valor com
    // `array[índice]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // A função interna `len` retorna o tamanho de um array.
    fmt.Println("len:", len(a))

    // Use esta sintaxe para declarar e inicializar um array
    // com uma linha.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Tipos de arrays são unidimensionais, mas você
    // pode compor mais tipos para construir estruturas
    // multidimensionais.
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
