// Podemos usar canais para sincronizar a execução entre
// gorrotinas. Aqui está um exemplo de como bloquear
// o recebimento para aguardar o termino de uma gorrotina.

package main

import "fmt"
import "time"

// Esta é a função que vamos executar em uma gorroutina.
// O canal `done` será usado para notificar outra gorrotina
// que o trabalho desta função está concluído.
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // Envie um valor para notificar o término.
    done <- true
}

func main() {

    // Inicia uma gorrotina `worker` passando para ela
    // o canal para notificação.
    done := make(chan bool, 1)
    go worker(done)

    // Aguarda o recebimento da notificação do canal em
    // worker.
    <-done

}
