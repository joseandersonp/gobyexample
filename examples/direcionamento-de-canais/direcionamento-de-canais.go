// Ao usar canais como parâmetros de função, você pode
// especificar se um canal serve apenas para enviar ou
// receber valores. Esta particularidade aumenta a
// segurança dos tipos do programa.
package main

import "fmt"

// A função `ping` só aceita um canal para enviar valores.
// Isso causaria um erro de compilação ao tentar receber por
// este canal.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// A função `pong` aceita um canal para receber (`pings`)
// e um segundo para enviar (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
