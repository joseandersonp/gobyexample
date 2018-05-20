// Freqüentemente queremos executar o código Go em algum
// momento no futuro, ou repetidamente com um intervalo de
// tempo. Os recursos integrados de _timer_ e _ticker_ do
// Go facilitam essas duas tarefas. Veremos primeiro
// os timers e logo em seguida os [tickers](tickers).

package main

import "time"
import "fmt"

func main() {

    // Os Timers representam um único evento no futuro.
    // Você informa ao timer quanto tempo deseja esperar
    // e fornece um canal que será notificado esse momento.
    // Este timer irá aguardar 2 segundos.
    timer1 := time.NewTimer(2 * time.Second)

    // O `<-timer1.C` faz uma pausa no canal do timer `C`
    // até que ele envie um valor indicando que o timer
    // expirou.
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // Se você quisesse apenas esperar, você poderia ter
    // usado `time.Sleep`. Um motivo pelo qual um timer
    // pode ser útil é que você pode cancelá-lo antes que
    // ele expire. Aqui está um exemplo disso.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
