# Autenticação e Autorização

A API utiliza JWT para autenticação e bcrypt para armazenamento seguro de senhas.

## Login

```http
POST /auth/login
Content-Type: application/json
```

```json
{
  "username": "admin",
  "password": "senha123"
}
```

Resposta:

```json
{
  "token": "<jwt>",
  "token_type": "Bearer",
  "expires_in": 86400,
  "user": {
    "id": 1,
    "name": "Administrador",
    "username": "admin",
    "email": "admin@biblioteca.com",
    "role": "admin"
  }
}
```

O token possui validade de 24 horas e é assinado com HS256.

## Uso do token

Envie o token em todas as rotas protegidas:

```http
Authorization: Bearer <jwt>
```

Tokens ausentes, inválidos ou expirados resultam em `401 Unauthorized`.

## Usuário autenticado

```http
GET /auth/me
Authorization: Bearer <jwt>
```

Retorna os dados do usuário autenticado.

## Perfis

O sistema possui os perfis:

- `admin`
- `user`

A autorização por perfil é validada pela função `security.ValidateRoles`, que recebe as claims do JWT e um ou mais roles permitidos.

Exemplo:

```go
security.ValidateRoles(
    claims,
    entity.RoleAdmin,
)
```

Quando o usuário não possui um role permitido, a ação é recusada.

## Segurança

- Senhas são armazenadas com bcrypt.
- O hash da senha não é retornado pela API.
- O middleware valida assinatura e expiração do JWT.
- Usuários inativos não podem realizar login.
- O username é único.
- O email é validado para evitar duplicidade.
