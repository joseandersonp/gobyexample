# O primeiro timer irá expirar ~2s depois de iniciarmos
# o programa, mas o segundo deve ser cancelado antes que
# ele tenha a chance de expirar.
$ go run timers.go
Timer 1 expired
Timer 2 stopped
