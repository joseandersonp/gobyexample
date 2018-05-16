// _Canais_ são os vias que conectam goroutines
// concorrentes. Você pode enviar valores para os
// canais de uma goroutine e receber esses valores
// em outra goroutine.

package main

import "fmt"

func main() {

	// Crie um novo canal com `make(chan tipo-valor)`.
	// Os canais são do tipo dos valores que eles
	// transmitem.
	messages := make(chan string)

	// _Envie_ um valor para um canal usando a sintaxe
	// `canal <-`. Aqui nós enviamos `"ping"` para o
	// canal `messages` que fizemos anteriormente de
	// outra nova goroutine.
	go func() { messages <- "ping" }()

	// A sintaxe `<-canal` recebe um valor do canal.
	// Aqui vamos receber a mensagem `"ping"` que
	// enviamos anteriormente e apresentamos na tela.
	msg := <-messages
	fmt.Println(msg)
}
