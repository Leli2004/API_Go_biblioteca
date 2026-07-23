
# API Biblioteca

Sistema de gerenciamento de biblioteca desenvolvido em Go.

A API tem como objetivo controlar livros, autores, editoras, usuários, empréstimos, devoluções, reservas, multas e fila de espera.

---

## Ferramentas
- Go v1.26
- PostgreSQL v16.14
- Docker v29.6
- Redis v7-alpine
- Make

---

## Como executar o projeto

### 1. Subir containeres

```bash
make db-up
make redis-up
````

O PostgreSQL será iniciado na porta `5433`.

### 2. Criar o banco da aplicação

```bash
make db-create
```

### 3. Executar a API

```bash
make run
```

---

## Comandos úteis

Lista todos os comandos disponíveis:
```bash
make help
```

Acessa o terminal do PostgreSQL dentro do container:
```bash
make db-shell
```

Exibe os logs do container PostgreSQL:
```bash
make db-logs
```

Remove o container do PostgreSQL:
```bash
make db-down
```

---

## Arquitetura do projeto

```txt
.
├── cmd
├── config
├── db
├── dist
├── internal
└── migrate
```

### Descrição das pastas

* `cmd`: ponto de entrada da aplicação, onde fica o `main.go`.
* `config`: configurações da aplicação, como variáveis de ambiente e conexão com banco.
* `db`: conexão e configurações relacionadas ao banco de dados.
* `dist`: arquivos gerados após o build da aplicação.
* `internal`: regras de negócio, handlers, repositories e entidades internas.
* `migrate`: arquivos de migration para criação e alteração das tabelas.

---

## Entidades principais

* `genre`: gênero do livro.
* `author`: autor do livro.
* `publisher`: editora responsável pela publicação.
* `book`: livro cadastrado no sistema.
* `user`: usuário da biblioteca.

---

## Ações principais

* `loan`: empréstimo de livro.
* `return`: devolução de livro.
* `reservation`: reserva de livro.
* `fine`: multa por atraso na devolução.
* `waitlist`: fila de espera para empréstimo.

---

## TODO

* [ ] Implementar cache com Redis.

* [ ] Front-end (somente para facilitar uso das requests, diretório ./web).

---
