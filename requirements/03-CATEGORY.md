# 🗂 Category

## RF-CAT-01 — Listar Categorias

Endpoint: GET */category*

A API deve retornar uma lista de categorias cadastradas.

## RF-CAT-02 — Criar Categoria

Endpoint: POST */category*

A API deve permitir a criação de uma nova categoria.
* Nome
* Status
* Usuário

Regras:
* Deve exigir autenticação
* Nome da categoria deve ser obrigatório

## RF-CAT-03 — Buscar Categoria por ID

Endpoint: GET */category/{id}*

A API deve retornar os dados de uma categoria específica.

## RF-CAT-04 — Atualizar Categoria

Endpoint: PUT */category/{id}*

A API deve permitir a atualização dos dados de uma categoria existente.

## RF-CAT-05 — Remover Categoria

Endpoint: DELETE */category/{id}*

A API deve permitir a remoção de uma categoria existente.