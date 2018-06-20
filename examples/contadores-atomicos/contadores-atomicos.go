// O principal mecanismo para gerenciamento de estados em Go
// é a comunicação através dos canais. Vimos isso com os
// [pools de trabalhadores](pool-de-trabalhadores). Existem
// outras opções para gerenciamento de estados. Aqui,
// analisaremos o uso do pacote `sync/atomic` para
// _contadores atômicos_ acessados ​​por várias gorroutinas.

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// Usaremos um inteiro sem sinal (sempre positivo)
	// para representar nosso contador .
	var ops uint64

	// Para simular simultâneamente as atualizações,
	// iniciamos cinquenta gorrotinas que incrementam o
	// contador uma vez por milissegundo.
	for i := 0; i < 50; i++ {
		go func() {
			for {

				// Para incrementar atomicamente o contador,
				// usamos o `AddUint64`, fornecendo-lhe o
				// endereço de memória do nosso contador
				// `ops` com a sintaxe `&`.
				atomic.AddUint64(&ops, 1)

				// Aguarde um pouco entre os incrementos.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Aguarde um segundo para permitir que alguns
	// ops se acumulem.
	time.Sleep(time.Second)

	// Para usar o contador com segurança enquanto ele
	// ainda está sendo atualizado por outras gorrotinas,
	// extraímos uma cópia do valor atual no `opsFinal`
	// por meio do `LoadUint64`. Como acima, precisamos
	// dar a esta função o endereço de memória `&ops`
	// do qual vamos pegar o valor.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
