// Neste exemplo, veremos como implementar um _pool de
// trabalhadores_ usando gorrotinas e canais.
package main

import "fmt"
import "time"

// Aqui está o trabalhador (worker), do qual vamos executar
// várias instâncias concorrentes. Esses trabalhadores receberão
// uma tarefa no canal `jobs` e enviarão os resultados
// correspondentes em `results`. Vamos usar um `time.Sleep` para
// fazer uma pausa de um segundo por tarefa para simular uma
// tarefa dispendiosa.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished", j)
        results <- j * 2
    }
}

func main() {

    // Para usar nosso pool de trabalhadores, precisamos
    // enviá-los para o trabalho e coletar seus resultados.
    // Criamos dois canais para isso.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Este loop inicia três trabalhadores, inicialmente
    // parados porque ainda não há tarefas a executar.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Aqui nós enviamos cinco tarefas para o canal `jobs`
    // e depois fechamos esse canal com `close` para indicar
    // que é todo o trabalho que temos.
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Finalmente coletamos todos os resultados do trabalho.
    for a := 1; a <= 5; a++ {
        <-results
    }
}
