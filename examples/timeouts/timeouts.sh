# Ao executar este programa podemos ver que o tempo limite
# da primeira operação expira e o segunda operação e 
# executada com sucesso.
$ go run timeouts.go 
timeout 1
result 2

# Usar esse padrão de `select` com timeout requer a 
# comunicação de resultados através de canais. Esta é 
# uma boa idéia em geral, porque outros recursos 
# importantes do Go são baseados em canais e `select`.
# Vamos ver dois exemplos disso: timers e tickers.