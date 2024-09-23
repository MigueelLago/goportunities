# GoVagas Vagas API

Este projeto é uma API moderna de vagas de emprego construída usando Golang, como forma de aprender mais sobre a linguagem. A API utiliza o Go-Gin como roteador, GoORM para comunicação com o banco de dados, SQLite como banco de dados e Swagger para documentação e testes da API. O projeto segue uma estrutura de pacotes moderna para manter a base de código organizada e de fácil manutenção.

---

## Features

- Introdução ao Golang e construção de APIs modernas
- Configuração do ambiente de desenvolvimento para criar a API
- Uso do Go-Gin como roteador para gerenciamento de rotas
- Implementação do SQLite como banco de dados para a API
- Uso do GoORM para comunicação com o banco de dados
- Integração do Swagger para documentação da API
- Criação de uma estrutura de pacotes moderna para organizar o projeto
- Implementação de uma API completa de vagas de emprego com endpoints para busca, criação, edição e exclusão de vagas

## Makefile Commands

O projeto inclui um Makefile para ajudá-lo a gerenciar tarefas comuns com mais facilidade. Aqui está uma lista dos comandos disponíveis e uma breve descrição do que eles fazem:

- `make docs`: Gera a documentação da API usando Swag.
- `make run`: Executa a aplicação.
- `make build`: Compile o aplicativo e crie um arquivo executável chamado `gopportunities`.

## Docker and Docker Compose

Este projeto inclui um arquivo `Dockerfile` e `docker-compose.yml` para fácil conteinerização e implantação. Aqui estão os comandos Docker e Docker Compose mais comuns que você pode querer usar:

- `docker compose build`: Construa os serviços definidos no arquivo `docker-compose.yml`.
- `docker compose up`: Execute os serviços definidos no arquivo `docker-compose.yml`.

Para parar e remover containers, redes e volumes definidos no arquivo `docker-compose.yml`, execute:

```sh
docker-compose down
```

Para obter mais informações sobre Docker e Docker Compose, consulte a documentação oficial:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Ferramentas Utilizadas

Este projeto utiliza as seguintes ferramentas:

- [Golang](https://golang.org/) para desenvolvimento de backend
- [Go-Gin](https://github.com/gin-gonic/gin) para gerenciamento de rotas
- [GoORM](https://gorm.io/) para comunicação com o banco de dados
- [SQLite](https://www.sqlite.org/index.html) banco de dados
- [Swagger](https://swagger.io/) para documentação da API

## Uso

Depois que a API estiver em execução, você poderá usar a UI do Swagger para interagir com os endpoints para pesquisar, criar, editar e excluir oportunidades de emprego. A API pode ser acessada em `http://localhost:$PORT/swagger/index.html`.

Padrão $PORT se não for fornecido: 8080.

## License

Este projeto está licenciado sob a licença MIT.

## Credits

Projeto criado com base na Live do Arthur: [arthur404dev](https://github.com/arthur404dev).