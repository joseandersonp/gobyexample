// Por padrão, os canais não usam buffer, o que significa
// que eles só aceitarão envios (`chan <-`) se houver um
// recebimento correspondente (`<- chan`) pronto para receber
// o valor enviado. Os canais com buffer aceitam um número
// limitado de valores sem um receptor correspondente para
// esses valores.

package main

import "fmt"

func main() {

    // Aqui usamos 'make' para criar um canal de strings para
    // armazenar dois 2 valores.
    messages := make(chan string, 2)

    // Como esse canal possui um buffer, podemos enviar estes
    // valores para o canal sem um recebimento simultâneo
    // correspondente.
    messages <- "buffered"
    messages <- "channel"

    // Como de costume, podemos receber estes dois valores
    // mais tarde.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
