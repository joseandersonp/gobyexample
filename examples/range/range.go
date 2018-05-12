// `range` itera sobre os elementos de uma variedade de
// estruturas de dados. Vamos ver como usar o `range` com
// algumas das estruturas de dados que já aprendemos.

package main

import "fmt"

func main() {

    // Aqui usamos `range` para somar os números de um slice.
    // Arrays funcionam da mesma maneira.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` em arrays e slices fornece o índice e o valor
    // de cada entrada. Anteriormente, não precisávamos do
    // índice, então o ignoramos com o identificador de
    // ausência `_` . Entretanto, em alguns casos vamos
    // precisar dos índices.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range` itera com os pares chave/valor de um mapa.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` também pode iterar apenas as chaves de um mapa.
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // `range` em strings itera sobre os pontos do código
    // Unicode. O primeiro valor é o byte de índice do `rune`
    // e o segundo a próprio `rune`.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
