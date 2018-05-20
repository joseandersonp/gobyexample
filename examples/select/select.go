// O _select_ do Go permite aguardar o resultado das
// várias operações de canal. Combinando gorrotinas
// e canais com select é um poderoso recurso do Go.

package main

import "time"
import "fmt"

func main() {

    // Para o nosso exemplo vamos utilizar select com
    // dois canais
    c1 := make(chan string)
    c2 := make(chan string)

    // Cada canal receberá um valor após um certo tempo,
    // para simular, por exemplo, bloqueios de operações
    // de RPC executando em gorrotinas concorrentes.
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // Usaremos o `select` para aguardar por dois valores
    // simultaneamente, apresentando na tela cada valor à
    // medida que ele chega.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
