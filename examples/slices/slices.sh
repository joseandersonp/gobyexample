# Note que, embora os slices sejam tipos diferentes dos
# arrays, eles são apresentados por `fmt.Println` de 
# forma parecida.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Confira este [excelente post do blog](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html) da equipe 
# do Go para obter mais detalhes sobre o design e 
# a implementação de slices em Go.

# Agora que aprendemos sobre arrays e slice, vamos ver
# outra estrutura de dados fundamental do Go: os mapas.
