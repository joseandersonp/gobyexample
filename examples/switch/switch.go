// _Instruções Switch_ expressam desvios condicionais
// atravéz de várias condições.

package main

import "fmt"
import "time"

func main() {

    // Aqui está um `switch` básico.
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // Você pode usar vírgulas para separar várias expressões
    // na mesma declaração `case`. Usamos o caso opcional
    // `default` neste exemplo também.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    // `switch` sem uma expressão é uma maneira alternativa
    // de expressar um if/else. Aqui também mostramos como
    // as expressões `case` podem ser não-constantes.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // Um `switch` para _tipos_ compara os tipos em vez dos
    // valores. Você pode usar isso para descobrir o tipo de
    // um valor de interface. Neste exemplo, a variável `t`
    // terá o tipo correspondente à sua cláusula.
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
