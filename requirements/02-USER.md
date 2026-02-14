# 👤 User
## RF-USER-01 — Criar Usuário

Endpoint: POST */user*

A API deve permitir a criação de um novo usuário, informando dados básicos como:
* Nome
* Email
* Senha

Critérios de Aceite:
* Validar Nome e Email
* O usuário deve ser persistido no banco de dados
* Não permitir emails duplicados
* Retornar 201 Created

## RF-USER-02 — Obter Dados do Usuário Autenticado

Endpoint: GET */user*

A API deve retornar os dados do usuário autenticado.

Regras:
* O endpoint deve exigir autenticação
* Retornar apenas os dados do próprio usuário

## RF-USER-03 — Atualizar Usuário

Endpoint: PUT */user*

A API deve permitir que o usuário autenticado atualize seus próprios dados.

Critérios de Aceite:
* Deve validar permissões
* Retornar dados atualizados
* Retornar 200 OK