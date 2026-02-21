# 📦 Product

## RF-PROD-01 — Listar Produtos

Endpoint: GET */product*

A API deve retornar uma lista de produtos cadastrados.

## RF-PROD-02 — Criar Produto

Endpoint: POST */product*

A API deve permitir a criação de um produto associado a uma categoria.
* Nome
* Descriçåo
* Preço
* Status
* Categoria

Regras:
* Deve exigir autenticação
* Nome da produto deve ser obrigatório

## RF-PROD-03 — Buscar Produto por ID

Endpoint: GET */product/{id}*

A API deve retornar os dados de um produto específico.

## RF-PROD-04 — Atualizar Produto

Endpoint: PUT */product/{id}*

A API deve permitir a atualização dos dados de um produto existente.

## RF-PROD-05 — Remover Produto

Endpoint: DELETE */product/{id}*

A API deve permitir a remoção de um produto existente.