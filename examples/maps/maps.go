// _Mapas_ são um [tipo de dados associativos] incorporado do Go
// (às vezes chamado de _hashes_ ou _dicionários_ em outras linguagens).

package main

import "fmt"

func main() {

    // Para criar um mapa vazio, use a função incorporada
    // `make`: `make(map[typo-chave]tipo-valor)`.
    m := make(map[string]int)

    // Defina os pares chave/valor usando a sintaxe típica
    // `nome[chave] = valor`.
    m["k1"] = 7
    m["k2"] = 13

    // Imprimindo um mapa com `fmt.Println`, será apresentado
    // todos os seus pares chave/valor.
    fmt.Println("map:", m)

    // Obtenha um valor de uma chave com `nome[chave]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // A função incorporada `len` retorna a quantidade de pares
    // quando chamada com um mapa.
    fmt.Println("len:", len(m))

    // A função incorporada `delete` remove um par chave/valor
    // de um mapa.
    delete(m, "k2")
    fmt.Println("map:", m)

    // O segundo valor de retorno opcional de um mapa indica se
    // a chave está presente no mapa. Isso pode ser usado para
    // destinguir entre chaves desconhecidas e chaves com valor-zero
    // como `0` ou `""`. Aqui não precisamos do valor em si, então o
    // ignoramos com o _identificador de ausência_ `_`.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // Você também pode declarar e inicializar um novo mapa
    // na mesma linha.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
