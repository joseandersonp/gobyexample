// Go suporta
// [_funções recursivas_](http://en.wikipedia.org/wiki/Recursion_(computer_science)).
// Aqui está um exemplo clássico de fatorial.

package main

import "fmt"

// Essa função `fact` chama a si própria até atingir a
// base `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
