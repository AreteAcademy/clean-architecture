# 📦 Product

## RF-PROD-01 — Criar Produto

Endpoint: POST */product*

A API deve permitir a criação de um produto associado a uma categoria e usuário.
* Nome
* Descriçåo
* Preço
* Status
* Categoria
* Usuário

Regras:
* Nome do produto deve ser obrigatório
* Descriçåo do produto deve ser obrigatório
* Preço do produto deve ser obrigatório
* Status do produto deve ser obrigatório
* Deve aceitar somente status *ACTIVE* e *INACTIVE*
* Categoria deve ser obrigatório
* Usuário deve ser obrigatório
* Deve exigir autenticação

## RF-PROD-02 — Listar Produtos

Endpoint: GET */product*

A API deve retornar uma lista de produtos cadastrados.

## RF-PROD-03 — Buscar Produto por ID

Endpoint: GET */product/{id}*

A API deve retornar os dados de um produto específico.

## RF-PROD-04 — Atualizar Produto

Endpoint: PUT */product/{id}*

A API deve permitir a atualização dos dados de um produto existente.

## RF-PROD-05 — Remover Produto

Endpoint: DELETE */product/{id}*

A API deve permitir a remoção de um produto existente.