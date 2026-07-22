
# Autenticação JWT

## Visão Geral

O sistema utiliza autenticação baseada em **JWT (JSON Web Token)**.

Após realizar o login com **username** e **password**, a API gera automaticamente um token JWT com validade de **24 horas**. Esse token deve ser enviado em todas as requisições para rotas protegidas.

---

# Fluxo de Autenticação

```text
Usuário
    │
    ▼
POST /auth/login
    │
    ▼
Validação de username e senha
    │
    ▼
Geração do JWT
    │
    ▼
Retorno do token
    │
    ▼
Authorization: Bearer <token>
    │
    ▼
Rotas protegidas
```

---

# Cadastro de Usuário

Durante o cadastro:

- A senha é recebida em texto puro.
- A senha é criptografada utilizando **bcrypt**.
- Apenas o **password_hash** é armazenado no banco de dados.
- A senha nunca é retornada pela API.

---

# Login

## Endpoint

```http
POST /auth/login
```

### Request

```json
{
    "username": "admin",
    "password": "senha123"
}
```

### Response

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

---

# Usuário Autenticado

## Endpoint

```http
GET /auth/me
```

### Header obrigatório

```http
Authorization: Bearer <jwt>
```

### Response

```json
{
    "id": 1,
    "name": "Administrador",
    "username": "admin",
    "email": "admin@biblioteca.com",
    "role": "admin"
}
```

---

# Header de Autenticação

Todas as rotas protegidas devem receber o seguinte header:

```http
Authorization: Bearer <jwt>
```

Caso o token esteja ausente, inválido ou expirado, a API retornará:

```http
401 Unauthorized
```

---

# Perfis

O sistema possui dois perfis de usuário:

| Perfil | Descrição   |
|---------|------------------------- |
| `admin` | Administrador do sistema |
| `user` | Usuário padrão  |

---

# Expiração do Token

- Algoritmo: **HS256**
- Validade: **24 horas**
- Após a expiração, o usuário deverá realizar um novo login para obter um novo token.

---

# Segurança

- Senhas armazenadas utilizando **bcrypt**.
- Apenas o **password_hash** é persistido no banco de dados.
- O JWT é assinado utilizando **HS256**.
- O middleware valida automaticamente a assinatura e a expiração do token.
- Senhas e hashes nunca são retornados nas respostas da API.
- Usuários inativos não podem realizar login.

---
