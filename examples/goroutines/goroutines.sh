# Quando executamos este programa, vemos a saída da primeira 
# chamada, depois a saída intercalada das duas goroutines. 
# Essa intercalação reflete as goroutines sendo executadas 
# simultaneamente em tempo de execução do Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# A seguir, veremos um complemento para goroutines
# em programas concorrentes em Go: channels.
