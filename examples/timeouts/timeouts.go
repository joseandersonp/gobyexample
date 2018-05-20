// _Timeouts_ são importantes para programas que se
// conectam a recursos externos ou que de outra
// forma precisam limitar o tempo de execução.
// A implementação de timeouts no Go é fácil e
// elegante graças aos canais e `select`.
package main

import "time"
import "fmt"

func main() {

    // Para o nosso exemplo, suponha que estamos
    // executando uma chamada externa que retorna seu
    // resultado em um canal `c1` depois de 2s.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()

    // Aqui temos o `select` implementando um timeout.
    //`res: = <-c1` aguarda o resultado e <-Time.Após aguarda
    // o valor que será enviado após o time out de 1s.
    // Conforme o `select` prossegue com a primeira mensagem
    // recebida, consideraremos o case de timeout se a operação
    // demorar mais do que o tempo permitido (1s).
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    // Se permitirmos um timeout de 3s, o recebimento de `c2`
    // será bem-sucedido e nós imprimiremos o resultado.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
