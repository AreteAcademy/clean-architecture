# 📄 Documento de Requisitos – API REST (Curso)

## 1. Visão Geral

Este documento descreve os requisitos funcionais da API REST desenvolvida para o curso, garantindo que os alunos compreendam, implementem e validem corretamente os endpoints propostos.

A API segue os princípios de:

* Arquitetura REST
* Separação de responsabilidades
* Autenticação via token
* Operações CRUD completas

## 2. Requisitos Gerais da API
### RF-GERAL-01 — Protocolo HTTP

A API deve operar exclusivamente via protocolo HTTP/HTTPS.

### RF-GERAL-02 — Formato de Dados

A API deve aceitar e retornar dados no formato JSON.

### RF-GERAL-03 — Códigos de Status HTTP

A API deve retornar códigos HTTP apropriados, incluindo:

* 200 OK
* 201 Created
* 400 Bad Request
* 401 Unauthorized
* 404 Not Found
* 500 Internal Server Error

### RF-GERAL-04 — Autenticação
Endpoints protegidos devem exigir autenticação, utilizando token JWT enviado via header Authorization.

## 3. Requisitos por Módulo
🔍 Health
### RF-HEALTH-01 — Verificação de Saúde
Endpoint: GET */health*

A API deve disponibilizar um endpoint de health check que:
* Retorne status HTTP 200
* Indique que a aplicação está operacional
* Não exija autenticação

Critério de Aceite:
```
{
  "status": "ok"
}
```