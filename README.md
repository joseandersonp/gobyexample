## GO por Exemplo

Conteúdo e ferramentas para construção do [Go por Exemplo](https://goporexemplo.com),
um site que ensina Go via programas de exemplo comentados.


### Visão Geral

Tradução do site Go por Exemplo, construído pela extração 
do código e comentários de arquivos fonte em `examples` e 
renderizando-os via `templates` para dentro do diretório 
estático `public`. Os programas que implementam este 
processo de construção estão em `tools`, junto com algumas 
dependências de fornecedores em `vendor`.

O diretório `public` pode ser hospedado por qualquer 
sistema de conteúdo estático. A produção do site 
usa S3 e CloudFront, por exemplo.


### Construindo

Para construir o site, você precisará instalar o Go 
e o Python. Execute:

```console
$ go get github.com/russross/blackfriday
$ tools/build
$ open public/index.html
```

Para construir continuamente em um loop:

```console
$ tools/build-loop
```

### Publicando

To hospedar o site:

```console
$ gem install aws-sdk
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload-site
```

### Licença

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Traduções

As traduções de colaboradores do site Go por exemplo estão disponíveis em:

* [Chinês] (https://gobyexample.xgwang.me/) por [xg-wang] (https://github.com/xg-wang/gobyexample)
* [Francês] (http://le-go-par-l-exemple.keiruaprod.fr) por [keirua] (https://github.com/keirua/gobyexample)
* [Italiano] (http://gobyexample.it) pela [Comunidade Go Italia] (https://github.com/golangit/gobyexample-it)
* [Coreano] (https://mingrammer.com/gobyexample/) por [mingrammer] (https://github.com/mingrammer)
* [Espanhol] (http://goconejemplos.com) pela [comunidade Go Mexico] (https://github.com/dabit/gobyexample)

### Agradecimentos

Thanks to [Jeremy Ashkenas](https://github.com/jashkenas)
for [Docco](http://jashkenas.github.com/docco/), which
inspired this project.
