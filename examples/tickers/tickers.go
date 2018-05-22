// [Timers](timers) são usados para fazer algo uma vez no
// futuro - _tickers_ são usados para fazer algo repetidamente
// em intervalos regulares de tempo. Aqui está um exemplo
// de um ticker que executa periodicamente até pararmos.

package main

import "time"
import "fmt"

func main() {

	// Tickers usam um mecanismo semelhante aos timers: um
	// um canal para o qual os valores são enviados. Aqui
	// usaremos o `range` no canal para iterar os valores
	// conforme eles chegam a cada 500 ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick em", t)
		}
	}()

	// Tickers podem ser parados assim como os timers. Depois
	// que um ticker é parado, ele não receberá mais valores em
	// seu canal. Vamos parar o nosso depois de 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker parado")
}
