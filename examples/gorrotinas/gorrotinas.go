// _Gorrotina_ é a execução de uma thread de processamento
// leve.

package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    // Suponha que temos uma chamada da função `f(s)`. Veja
    // como chamaríamos da maneira normal, executando a de
    // forma síncrona.
    f("direto")

    // Para chamar a função em uma gorrotina, usamos
    // `go f(s)`. Esta nova gorrotina será executada
    // simultaneamente com quem está chamando.
    go f("gorrotina")

    // Você também pode iniciar uma gorrotina para uma
    // chamada de função anônima.
    go func(msg string) {
        fmt.Println(msg)
    }("indo")

    // Nossas chamadas de função estão sendo executadas de
    // forma assíncrona em duas gorrotinas separadas, então
    // a execução avança até aqui. Este `Scanln` requer
    // que pressionemos uma tecla antes do programa sair.
    fmt.Scanln()
    fmt.Println("fim")
}
