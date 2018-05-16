// _Interfaces_ são coleções nomeadas de assinaturas
// de métodos.

package main

import "fmt"
import "math"

// Aqui está uma interface básica para formas
// geométricas.
type geometry interface {
	area() float64
	perim() float64
}

// Para o nosso exemplo, vamos implementar essa interface
// nos tipos `rect` e` circle`.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.

// Para implementar uma interface no Go, só precisamos
// implementar todos os métodos da interface. Aqui nós
// implementamos `geometry` em `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// A implementação para os `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Se uma variável possuir um tipo interface, podemos chamar
// métodos que estão na interface nomeada. Aqui está uma
// função de `measure` genérica que aproveita isso para
// trabalhar em qualquer `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Os tipos struct `circle` e` rect` implementam a
	// interface `geometry` para que possamos usar
	// instâncias dessas estruturas como argumentos para
	// `measure`.
	measure(r)
	measure(c)
}
