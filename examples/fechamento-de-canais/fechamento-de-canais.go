// _Fechar_ um canal indica que os valores não serão
// mais enviados por este. Isso pode ser útil para
// comunicar a conclusão aos canais receptores.

package main

import "fmt"

// Neste exemplo, vamos usar o canal `jobs` para comunicar
// o trabalho a ser executado a partir da gorrotina `main()`
// para a gorrotina _trabalhador_. Quando não houver mais
// trabalhos (jobs) para trabalhador, fecharemos o canal
// `jobs` com a função incorporada `close`.
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // Aqui está a gorrotina trabalhador. Ela recebe continuamente
    // de `jobs` com `j, more: = <-jobs`. Nesta forma de recepção
    // de dois valores, o valor `more` será `false` se `jobs` tiver
    // sido fechado com `close` e todos os valores já tiverem sido
    // recebidos pelo canal. Usamos isso para notificar em `done`
    // o encerramento de todos os nossos trabalhos.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // Enviamos três trabalhos para o trabalhador através do canal
    // `jobs`, depois o fechamos com `close`.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // Aguardamos a finalização do trabalhador usando a abordagem
    // de sincronização(channel-synchronization) que vimos
    // anteriormente.
    <-done
}
