// _Funções_ são a parte central do Go. Aprenderemos
// sobre funções com alguns exemplos a seguir.

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.

// Está função que recebe dois `int`s e retorna o
// resultado da soma como um `int`.
func plus(a int, b int) int {

    //Go requer retornos explícitos, ou seja, não
    // retornará automaticamente o valor da última
    // expressão.
    return a + b
}

// Quando você passa vários parâmetros consecutivos do
// mesmo tipo, você pode omitir o tipo do nomes e
// declarar o tipo no ultimo parâmetro.
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    // Chame uma função como esperado, com
    // `nome(argumentos)`.
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
