// Desviar condicionalmente com `if` e `else`
// em Go é simples e rápido.

package main

import "fmt"

func main() {

    // Aqui está um exemplo básico.
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // Você pode ter uma instrução `if` sem um `else`.
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // Um desvio condicional pode preceder outros;
    // quaisquer variáveis ​​declaradas nesta declaração estão
    // disponíveis em todos os desvios condicionais.
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// Note que não é necessário parênteses em nas condições
// em Go, mas que as chaves são necessárias.
