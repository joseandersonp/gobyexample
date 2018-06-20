// A <em>[limitação de taxa](http://en.wikipedia.org/wiki/Rate_limiting)<em>
// é um mecanismo importante para controlar a utilização
// de recursos e manter a qualidade do serviço. Go suporta
// limitação taxa elegantemente com gorrotinas, canais e
// tickers.

package main

import "time"
import "fmt"

func main() {

	// Primeiro, vamos ver a limitação básica de taxa.
	// Suponha que queremos limitar nosso tratamento de
	// solicitações recebidas. Vamos atender a essas
	// solicitações em um canal com o mesmo nome.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// O canal `limiter` receberá um valor a cada
	// duzentos milissegundos. É o regulador no nosso
	// esquema de limitação de taxa.
	limiter := time.Tick(200 * time.Millisecond)

	// Ao bloquear durante a recepção do canal `limiter`
	// antes de atender a cada requisição, limitamo-nos
	// a uma requisição a cada duzentos milissegundos.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Podemos permitir pequenos picos de requisições em
	// nosso esquema de limitação de taxa, preservando o
	// limite geral da taxa. Podemos conseguir isso
	// bufferizando nosso canal limitador. O canal
	// `burstyLimiter` permitirá picos de até 3 eventos.
	burstyLimiter := make(chan time.Time, 3)

	// Abastecemos o canal para representar os picos
	// permitidos.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada duzentos milissegundos vamos tentar adicionar
	// um novo valor ao `burstyLimiter`, até o limite de três.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Agora simulamos mais cinco solicitações recebidas.
	// Os três primeiros se beneficiarão da capacidade de
	// pico do `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
