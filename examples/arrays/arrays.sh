# Note que que os arrays aparecem no formato
# `[v1 v2 v3 ...]` quando impressos com `fmt.Println`.
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

# Tipicamente você verá _slices_ com muito mais frequência
# do que arrays em Go. Vamos ver _slices_ a seguir.