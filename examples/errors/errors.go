// Em Go, é idiomático comunicar erros por meio de um
// explícito, valor de retorno separado. Isso contrasta com
// as exceções usadas em linguagens como Java e Ruby e o
// uníco valor sobrecarregado de resultado / erro às vezes
// usado em C. A abordagem de Go facilita a visualização de
// quais funções retornam erros e os manipulm usando as
// mesmas construções de linguagem empregadas para quaisquer
// outras tarefas que não sejam de erro.

package main

import "errors"
import "fmt"

// Por convenção, os erros são o últimos valores de
// retorno e possuem o tipo `error`, uma interface
// incorporada do Go.
func f1(arg int) (int, error) {
	if arg == 42 {

		//`errors.New` gera um valor de erro básico com a
		// mensagem de erro fornecida.
		return -1, errors.New("can't work with 42")

	}

	// Um valor `nil` na posição do erro indica que não
	// houve erro.
	return arg + 3, nil
}

// É possível usar tipos customizados como `error`s pela
// implementação do método `Error()`. Aqui temos uma
// variante no exemplo anterior que usa um tipo personalizado
// para representar explicitamente um argumento de erro.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// Neste caso, usamos a sintaxe `&argError` para
		// criar uma nova estrutura, fornecendo valores
		// para os dois campos `arg` e `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Os dois loops abaixo testam cada uma das nossas
	// funções de retorno de erros. Note que o uso de uma
	// verificação de erro em linha na declaração `if` é
	// comum na idiomática do código do Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Se você quiser usar programaticamente os dados em um
	// erro personalizado, você precisará obter o erro como
	// uma instância do tipo de erro personalizado por meio
	// de declaração de tipo.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
