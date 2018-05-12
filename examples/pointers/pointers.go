// Go suporta [_ponteiros_](http://en.wikipedia.org/wiki/Pointer_(computer_programming)),
// permitindo que você passe referências a valores e
// registros dentro do seu programa.

package main

import "fmt"

// Vamos ver como os ponteiros funcionam em contraste com os
// valores diretos usando duas funções: `zeroval` e `zeroptr`.
// `zeroval` recebe um parâmetro `int`, então os argumentos
// serão passados por valor. O `zeroval` receberá uma cópia
// do `ival` diferente da função onde é chamado.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.

// O `zeroptr` recebe um parâmetro `*int`, o que significa
// que ele recebe um ponteiro de `int`. O código `*iptr` no corpo
// da função _desreferencia_ o ponteiro de seu endereço de
// memória para o valor atual desse endereço. Se atribuirmos
// um valor a um ponteiro não referenciado, o valor que está
// sendo armazenado nesse endereço de memória é alterado.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// A sintaxe `&i` fornece o endereço de memória de `i`,
	// isto é, um ponteiro para `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Os ponteiros também podem ser impressos
	fmt.Println("pointer:", &i)
}
