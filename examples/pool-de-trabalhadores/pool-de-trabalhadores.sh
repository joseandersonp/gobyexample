# Nosso programa em execução mostra os cinco tarefas
# sendo executadas por vários trabalhadores. O programa 
# leva apenas cerca de dois segundos, apesar de fazer 
# cerca de cinco segundos de trabalho total, porque há
# três trabalhadores operando simultaneamente.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
