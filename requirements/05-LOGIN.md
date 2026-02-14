# 🔐 Login
## RF-AUTH-01 — Autenticação de Usuário

Endpoint: POST */login*

A API deve autenticar um usuário válido e retornar um token JWT.

Critérios de Aceite:
* Validar credenciais
* Retornar token JWT válido
* Retornar 401 em caso de falha

Exemplo de Resposta:
```
{
  "token": "jwt-token"
}
```