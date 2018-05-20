# Quando executamos o programa, a mensagem `"ping"` é
# passada com sucesso de uma goroutine para outra
# através do nosso canal.
$ go run channels.go 
ping

# Por padrão, a envios e as recebimentos são bloqueadas 
# até que o receptor e o transmissor estejam prontos. 
# Esta propriedade nos permite esperar até o final do
# programa pela mensagem "ping" sem ter que usar 
# qualquer outro tipo de sincronização.
