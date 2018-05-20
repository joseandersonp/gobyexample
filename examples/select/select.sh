# Como esperado, recebemos os valores `"one"` e 
# depois `"two"` 
$ time go run select.go 
received one
received two

# Note que o tempo de execução total é de apenas ~2 
# segundos porque ambos `Sleeps` executam de forma
# concorrente.
real	0m2.245s
