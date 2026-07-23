# Rotas da API

Todas as rotas abaixo exigem um JWT válido no header:

```http
Authorization: Bearer <jwt>
```

A única rota pública é `POST /auth/login`.

---

## Autenticação

| Método | Rota | Descrição |
|---|---|---|
| POST | `/auth/login` | Autentica um usuário com username e senha e retorna um JWT. |
| GET | `/auth/me` | Retorna os dados do usuário autenticado. |

---

## Gêneros

| Método | Rota | Descrição |
|---|---|---|
| GET | `/genre/list` | Lista os gêneros cadastrados. |
| GET | `/genre/:id` | Consulta um gênero pelo ID. |
| POST | `/genre` | Cadastra um gênero. |
| PUT | `/genre/:id` | Atualiza um gênero. |
| DELETE | `/genre/:id` | Remove um gênero. |

---

## Autores

| Método | Rota | Descrição |
|---|---|---|
| GET | `/author/list` | Lista os autores cadastrados. |
| GET | `/author/:id` | Consulta um autor pelo ID. |
| POST | `/author` | Cadastra um autor. |
| PUT | `/author/:id` | Atualiza um autor. |
| DELETE | `/author/:id` | Remove um autor. |

---

## Editoras

| Método | Rota | Descrição |
|---|---|---|
| GET | `/publisher/list` | Lista as editoras cadastradas. |
| GET | `/publisher/:id` | Consulta uma editora pelo ID. |
| POST | `/publisher` | Cadastra uma editora. |
| PUT | `/publisher/:id` | Atualiza uma editora. |
| DELETE | `/publisher/:id` | Remove uma editora. |

---

## Usuários

| Método | Rota | Descrição |
|---|---|---|
| GET | `/user/list` | Lista os usuários cadastrados. |
| GET | `/user/:id` | Consulta um usuário pelo ID. |
| POST | `/user` | Cadastra um usuário. |
| PUT | `/user/:id` | Atualiza um usuário. |
| DELETE | `/user/:id` | Remove um usuário. |

---

## Livros

| Método | Rota | Descrição |
|---|---|---|
| GET | `/book/list` | Lista os livros cadastrados. |
| GET | `/book/:id` | Consulta um livro pelo ID. |
| POST | `/book` | Cadastra um livro. |
| PUT | `/book/:id` | Atualiza um livro. |
| DELETE | `/book/:id` | Remove um livro. |

---

## Exemplares

| Método | Rota | Descrição |
|---|---|---|
| GET | `/book_copie/list` | Lista os exemplares cadastrados. |
| GET | `/book_copie/:id` | Consulta um exemplar pelo ID. |
| POST | `/book_copie` | Cadastra um exemplar. |
| PUT | `/book_copie/:id` | Atualiza um exemplar. |
| DELETE | `/book_copie/:id` | Remove um exemplar. |

---

## Empréstimos

| Método | Rota | Descrição |
|---|---|---|
| POST | `/loan/create` | Cria um empréstimo. |
| POST | `/loan/return` | Registra a devolução de um empréstimo. |
| GET | `/loan/list` | Lista os empréstimos. |
| GET | `/loan/:id` | Consulta um empréstimo pelo ID. |
| DELETE | `/loan/:id` | Remove um empréstimo. |

---

## Reservas

| Método | Rota | Descrição |
|---|---|---|
| POST | `/reservation/create` | Cria uma reserva de livro. |
