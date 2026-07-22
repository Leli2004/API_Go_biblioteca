# Rotas da API

---

## Auth

| Método | Rota | Descrição |
|---------|------|-----------|
| POST | `/auth/login` | Executa login, solicita username e senha. |
| GET | `/auth/me` | Retorna dados do usuário logado na sessão. |

---
## Gêneros

| Método | Rota | Descrição |
|---------|------|-----------|
| GET | `/genre/list` | Lista todos os gêneros cadastrados. |
| GET | `/genre/:id` | Consulta um gênero pelo seu identificador. |
| POST | `/genre` | Cadastra um novo gênero literário. |
| PUT | `/genre/:id` | Atualiza os dados de um gênero existente. |
| DELETE | `/genre/:id` | Remove um gênero do sistema. |

---

## Autores

| Método | Rota | Descrição |
|---------|------|-----------|
| GET | `/author/list` | Lista todos os autores cadastrados. |
| GET | `/author/:id` | Consulta um autor pelo seu identificador. |
| POST | `/author` | Cadastra um novo autor. |
| PUT | `/author/:id` | Atualiza os dados de um autor existente. |
| DELETE | `/author/:id` | Remove um autor do sistema. |

---

## Editoras

| Método | Rota | Descrição |
|---------|------|-----------|
| GET | `/publisher/list` | Lista todas as editoras cadastradas. |
| GET | `/publisher/:id` | Consulta uma editora pelo seu identificador. |
| POST | `/publisher` | Cadastra uma nova editora. |
| PUT | `/publisher/:id` | Atualiza os dados de uma editora existente. |
| DELETE | `/publisher/:id` | Remove uma editora do sistema. |

---

## Usuários

| Método | Rota | Descrição |
|---------|------|-----------|
| GET | `/user/list` | Lista todos os usuários cadastrados. |
| GET | `/user/:id` | Consulta um usuário pelo seu identificador. |
| POST | `/user` | Cadastra um novo usuário. |
| PUT | `/user/:id` | Atualiza os dados de um usuário existente. |
| DELETE | `/user/:id` | Remove um usuário do sistema. |

---

## Livros

| Método | Rota | Descrição |
|---------|------|-----------|
| GET | `/book/list` | Lista todos os livros cadastrados. |
| GET | `/book/:id` | Consulta um livro pelo seu identificador. |
| POST | `/book` | Cadastra um novo livro. |
| PUT | `/book/:id` | Atualiza os dados de um livro existente. |
| DELETE | `/book/:id` | Remove um livro do sistema. |

---

## Exemplares

| Método | Rota | Descrição |
|---------|------|-----------|
| GET | `/book_copie/list` | Lista todos os exemplares cadastrados. |
| GET | `/book_copie/:id` | Consulta um exemplar pelo seu identificador. |
| POST | `/book_copie` | Cadastra um novo exemplar de livro. |
| PUT | `/book_copie/:id` | Atualiza os dados de um exemplar existente. |
| DELETE | `/book_copie/:id` | Remove um exemplar do sistema. |

---

## Empréstimos

| Método | Rota | Descrição |
|---------|------|-----------|
| POST | `/loan/create` | Cria um novo empréstimo de um exemplar para um usuário. |
| POST | `/loan/return` | Registra a devolução de um empréstimo. |
| GET | `/loan/list` | Lista todos os empréstimos cadastrados. |
| GET | `/loan/:id` | Consulta um empréstimo pelo seu identificador. |
| DELETE | `/loan/:id` | Remove um empréstimo do sistema. |

---

## Reservas

| Método | Rota | Descrição |
|---------|------|-----------|
| POST | `/reservation/create` | Cria uma nova reserva para um livro indisponível. |
