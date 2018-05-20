// Os _structs_ do Go são coleções de campos com tipos
// específicos. Eles são úteis para agrupar dado e
// formar registros.

package main

import "fmt"

// Esse struct do tipo `person` possui os camos `name`
// e `age`
type person struct {
    name string
    age  int
}

func main() {

    // Essa sintaxe cria um novo struct.
    fmt.Println(person{"Bob", 20})

    // Você pode nomear os campos ao inicializar um struct.
    fmt.Println(person{name: "Alice", age: 30})

    // Os campos omitidos terão valor-zero.
    fmt.Println(person{name: "Fred"})

    // Um prefixo `&` produz um ponteiro para o struct.
    fmt.Println(&person{name: "Ann", age: 40})

    // Acesse os campos do struct com um ponto.
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // Você também pode usar pontos com os ponteiros
    // de structs - os ponteiros são automaticamente
    // desreferenciados.
    sp := &s
    fmt.Println(sp.age)

    // Structs são mutáveis.
    sp.age = 51
    fmt.Println(sp.age)
}
