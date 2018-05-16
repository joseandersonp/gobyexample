// Go permite definir _métodos_ para os structs.

package main

import "fmt"

type rect struct {
	width, height int
}

// Este método `area` possui um _receptor de tipo_ *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Os métodos podem ser definidos para os receptores de
// tipo de ponteiro ou valor. Aqui está um exemplo de um
// receptor de valor.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Aqui chamamos os dois métodos definidos para nosso
	// struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.

	// Go manipula automaticamente a conversão entre valores
	// e ponteiros para chamadas de método. Você pode querer
	// usar um receptor de tipo de ponteiro para evitar a
	// cópia em chamadas de método ou para permitir que o
	// método altere o struct receptor.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
