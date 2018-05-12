// _Slice_ são um tipo de dados Go que fornecem uma 
// interface mais poderosa para as sequências do que
// os arrays.

package main

import "fmt"

func main() {

    // Ao contrário dos arrays, os slices são definidos apenas 
    // pelos elementos que eles contêm (não pelo número de elementos). 
    // Para criar um slice vazio com comprimento diferente de zero,
    // use a funcão interna `make`. Aqui nós criamos um slice de `string`s 
    // com comprimento de `3` (inicializado com valor-zero).
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // Podemos definir e obter valores assim como nos arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])


    // Como esperado, `len` retorna o tamanho do slice.
    fmt.Println("len:", len(s))

    // Além dessas operações básicas, os slices suportam 
    // mais vários outras que os tornam mais ricos que os 
    // arrays. Uma delas é a função interna `append`, que
    // retorna um novos slice contendo um ou mais valores.
    // Note que precisamos atribuir um valor do retorno
    // `append`, pois obtemos um novo valor do tipo slice.
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Slices também podem ser copiados com `copy`. Aqui 
    // criamos um slice vazio `c` do mesmo tamanho que `s`
    // e copiamos de `s` para `c`.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Slices suportam o operador "slice" com a sintaxe 
    // `slice [low: high]`. Por exemplo, `l` obtém um slice
    // dos elementos `s[2]`, `s[3]` e `s[4]`.
    l := s[2:5]
    fmt.Println("sl1:", l)

    
    // Este retorna um slice até (mas excluindo) `s[5]`.
    l = s[:5]
    fmt.Println("sl2:", l)

    // e este returna um slice de (e incluíndo) `s[2]`.
    l = s[2:]
    fmt.Println("sl3:", l)    

    // Podemos declarar e inicializar uma variável para o 
    // slice em uma única linha apenas.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Slice podem ser compostos por estruturas de dados 
    // multidimensionais. O tamanho dos slices internos pode
    // variar, ao contrário dos arrays multidimensionais.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
