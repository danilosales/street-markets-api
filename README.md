# Street Markets API

A arquitetura do projeto foi montado utilizando as seguintes tecnologias:

* Go
* Gin Framework - Boilerplate para facilitar a construção dos endpoints REST
* Gorm - ORM em GO para persistência dos dados
* Testify - Para facilitar a escrita dos testes
* Zerolog - Biblioteca para logging
* Postgres - Base de dados para persistência das feiras livres.
* Swagger - Usado para docuemntar os serviços

## Instruções para executar o projeto


### Iniciar projeto

Após realizar o download do projeto, executar o seguinte comando:

```sh
 make run
```
Este comando acima irá subir alguns containers contendo a aplicação e o banco de dados, não sendo necessário mais nenhuma ação


### Iniciar projeto modo dev

Para iniciar em mode dev basta chamar o comando:

```sh
make dev
```

Este comando irá subir as depedencias para aplicação, em seguida pode executar o comando:

```sh
make run-dev
```

Endereços disponíveis:

* [Documentação dos serviços](http://localhost:8080/api/v1/swagger/index.html)

* [URL Base para os serviços](http://localhost:8080/api/v1/)

* [PGAdmin](http://localhost:54321) - User: admin@teste.com, Password: 123456

### Testes

Para executar os testes execute o comando:

```sh
make test
```

O seguinte comando abaixo gera um relatório de coverage dos testes:

```sh
make test-coverage
```

### Exemplos de Requisições


* **Cadastrar uma nova Feira Livre**

POST http://localhost:8080/api/v1/street-markets

```
{
    "long": "-46550164",
    "lat": "-23558733",
    "setcens": "355030885000091",
    "areap": "3550308005040",
    "coddist": 87,
    "distrito": "VILA FORMOSA",
    "codsubpref": 26,
    "subprefe": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "teste",
    "registro": "4041-2",
    "logradouro": "RUA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA"
}
```

* **Atualizar uma Feira Livre**

PUT http://localhost:8080/api/v1/street-markets/**4041-2** - Código de registro da feira

```
{
    "long": "-46550164",
    "lat": "-23558733",
    "setcens": "355030885000091",
    "areap": "3550308005040",
    "coddist": 87,
    "distrito": "VILA FORMOSA",
    "codsubpref": 26,
    "subprefe": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "teste",
    "registro": "4041-2",
    "logradouro": "RUA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA"
}
```

* **Excluir uma Feira Livre**

DELETE http://localhost:8080/api/v1/street-markets/**4041-2** - Código de registro da feira

* **Consultar uma Feira Livre**

GET http://localhost:8080/api/v1/street-markets/**4041-2** - Código de registro da feira

* **Buscar feiras**

As buscas por feiras podem ser feitas através de Query Params, tendo as seguintes possibilidades:

 - distrito
 - regiao
 - nome
 - bairro

Ex.:

GET http://localhost:8080/api/v1/street-markets?regiao=Leste

As buscas são realizadas de forma exata, e podem ser combinadas.

## Considerações Finais

* Acabei optando por utilizar um ORM por mera conveniência, mas para o atual estado do projeto sendo um CRUD sem muitas regras de negócio poderia ter utilizado as intruções SQL diretamente também.

* A importação dos dados foi feita diretamente ao banco, o correto seria fazer um tratamento mais adequado aos dados da planilha, normalizando a grafia por exemplo, e criando algumas entidades separadas como as informações de endereço.

* Os campos do endpoint de busca poderiam ser modificados para que a busca seja feita com like, mas visando uma melhor performance e assertividade o ideal seja utilizar algum motor de busca textual.

* Os logs da aplicação estão sendo gravados em um arquivo chamado street-markets.log, porém não foi criado nenhuma politica de rotação destes logs, mas os dados são gravados em json para que possa ser utilizado em algum agregador de log.


