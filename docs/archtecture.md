# Arquitetura do Projeto

O projeto é organizado por módulos de domínio, seguindo separação entre delivery HTTP, usecase, repository e entidades.

## Módulos principais

- `author`: gerenciamento de autores.
- `book`: gerenciamento de livros.
- `book_copie`: gerenciamento de exemplares.
- `genre`: gerenciamento de gêneros.
- `publisher`: gerenciamento de editoras.
- `user`: gerenciamento de usuários e credenciais.
- `auth`: login e autenticação JWT.
- `loan`: gerenciamento de empréstimos.
- `reservation`: gerenciamento de reservas.
- `fine`: processamento de multas.

## Autenticação

- `security`: geração, validação e parsing de JWT e comparação de senhas.
- `middleware`: valida o JWT nas requisições e disponibiliza `AuthClaims` para os handlers.
- `auth`: reutiliza o `user.Repository` para consultar usuários e executar o login.

## Autorização

A autenticação é feita pelo middleware JWT. A autorização por perfil é feita nos usecases que exigem restrição de acesso, utilizando `security.ValidateRoles`.

## Worker de multas

O worker de multas executa uma verificação periódica dos empréstimos vencidos e processa as multas conforme a regra definida no módulo `fine`.
