// A _goroutine_ é a execução de uma thread de
// processamento leve.

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suponha que temos uma chamada da função `f(s)`. Veja
	// como chamaríamos da maneira normal, executando a de
	// forma síncrona.
	f("direct")

	// Para chamar a função em uma goroutine, usamos
	// `go f(s)`. Esta nova goroutine será executada
	// simultaneamente com quem está chamando.
	go f("goroutine")

	// Você também pode iniciar uma goroutine para uma
	// chamada de função anônima.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Nossas chamadas de função estão sendo executadas de
	// forma assíncrona em duas goroutines separadas, então
	// a execução avança até aqui. Este `Scanln` requer
	// que pressionemos uma tecla antes do programa sair.
	fmt.Scanln()
	fmt.Println("done")
}
