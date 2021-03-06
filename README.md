## Go por Exemplo

Conteúdo e ferramentas para construção do [Go por Exemplo](https://goporexemplo.br),
um site que ensina Go via programas de exemplos comentados.


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

Para hospedar o site:

```console
$ gem install aws-sdk
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload-site
```

### Licença

Os direitos dessa tradução pertencem a José Anderson Pereira
e está sob licença de [Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Os direitos da obra original pertencem a Mark McGranaghan 
e está sob licença de [Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Os direitos de Go Gopher pertencem a [Renée French](http://reneefrench.blogspot.com/) e está sob licença de
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Traduções

As traduções de colaboradores do site Go por Exemplo estão disponíveis em:

* [Chinês](https://gobyexample.xgwang.me/) por [xg-wang](https://github.com/xg-wang/gobyexample)
* [Francês](http://le-go-par-l-exemple.keiruaprod.fr) por [keirua](https://github.com/keirua/gobyexample)
* [Italiano](http://gobyexample.it) pela [Comunidade Go Italia](https://github.com/golangit/gobyexample-it)
* [Coreano](https://mingrammer.com/gobyexample/) por [mingrammer](https://github.com/mingrammer)
* [Espanhol](http://goconejemplos.com) pela [comunidade Go Mexico](https://github.com/dabit/gobyexample)

### Agradecimentos

Agradecimentos a [Jeremy Ashkenas](https://github.com/jashkenas)
pelo [Docco](http://jashkenas.github.com/docco/), que
inspirou este projeto.
