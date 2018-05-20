// Envios e recebimentos básicos nos canais são bloqueados.
// No entanto, podemos usar usar `select` com uma condição
// `default` para implementar envios, recebimentos, e até
// mesmo `select`s multidirecionais _sem bloqueio_.
package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // Aqui está um recebimento sem bloqueio. Se um valor
    // estiver disponível em `messages`, então `select`
    // irá executar o case `<-messages` com esse valor.
    // Se não, ele executará imediatamente o case `default`.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // Um envio sem bloqueio funciona de maneira semelhante.
    // Aqui o `msg` não pode ser enviado para o canal
    // `messages`, porque o canal não tem buffer e não há
    // receptor. Portanto, o case `default` é executado.
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // Podemos usar vários `cases` acima de `default` para
    // implementar uma select multidirecional sem bloqueio.
    // Aqui tentamos recebimentos sem bloqueio tanto em
    // `messages` quanto em `signals`.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
