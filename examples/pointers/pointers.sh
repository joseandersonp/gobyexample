# `zeroval` não altera o `i` em `main`, mas 
# `zeroptr` sim porque tem uma referência
# do endereço de memória dessa variável.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
