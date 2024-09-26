# 🎬 API de Filmes

## 📝 Descrição Geral

Esta API permite o gerenciamento completo de filmes, incluindo criação, atualização, listagem e exclusão de registros em um banco de dados PostgreSQL. Desenvolvida em Go com o framework Gin, a API conta com autenticação básica e está dockerizada para facilitar o uso e implantação.

## 🏗️ Estrutura do Projeto

- **📂 Controller**: Gerencia as requisições HTTP e envia as respostas.
- **📂 Use Case**: Implementa a lógica de negócios e validações.
- **📂 Repository**: Responsável pelo acesso e manipulação dos dados no banco.
- **📂 Middleware**: Implementa a autenticação básica.
- **📂 DB**: Contém as configurações e gerencia a conexão com o PostgreSQL.
- **📂 Model**: Define as estruturas de dados utilizadas na aplicação.
- **📂 Routes**: Define as rotas de para conexão na aplicação.

## 🛣️ Endpoints

### 🟢 GET /ping

**Descrição**: Verifica se a API está online.

**Resposta de Sucesso**:
```json
{
  "message": "pong"
}
```

### 🟢 GET /filmes

**Descrição**: Retorna todos os filmes cadastrados.

**Exemplo de Resposta**:
```json
[
  {
    "id_filme": 1,
    "titulo": "O Poderoso Chefão",
    "diretor": "Francis Ford Coppola",
    "ano_lancamento": 1972,
    "genero": "Crime",
    "duracao": 175
  }
]
```

### 🟢 GET /filmes/{id}

**Descrição**: Retorna um filme específico pelo seu ID.

**Parâmetros de URL**:
- `id`: Identificador único do filme.

**Exemplo de Resposta**:
```json
{
  "id_filme": 1,
  "titulo": "O Poderoso Chefão",
  "diretor": "Francis Ford Coppola",
  "ano_lancamento": 1972,
  "genero": "Crime",
  "duracao": 175
}
```

### 🟢 POST /filmes

**Descrição**: Cria um novo filme.

**Corpo da Requisição**:
```json
{
  "titulo": "O Poderoso Chefão",
  "diretor": "Francis Ford Coppola",
  "ano_lancamento": 1972,
  "genero": "Crime",
  "duracao": 175
}
```

### 🟢 PUT /filmes/{id}

**Descrição**: Atualiza um filme existente.

**Parâmetros de URL**:
- `id`: Identificador único do filme.

**Corpo da Requisição**:
```json
{
  "titulo": "O Poderoso Chefão",
  "diretor": "Francis Ford Coppola",
  "ano_lancamento": 1972,
  "genero": "Crime",
  "duracao": 175
}
```

### 🟢 DELETE /filmes/{id}

**Descrição**: Remove um filme pelo ID.

**Parâmetros de URL**:
- `id`: Identificador único do filme.

## 🔐 Autenticação

A API utiliza autenticação básica em todas as rotas de filmes.

**Credenciais**:
- Usuário: `admin`
- Senha: `123`

## 🗄️ Banco de Dados

A conexão ao banco de dados PostgreSQL é configurada com as seguintes credenciais:

- **Host**: go_db
- **Port**: 5432
- **User**: postgres
- **Password**: admin
- **Database**: postgres

## 🚀 Como Rodar a API

### Pré-requisitos:
- Docker
- Docker Compose

### Passos:

1. Clone o repositório:
   ```
   git clone [URL_DO_REPOSITORIO]
   cd [NOME_DO_DIRETORIO]
   ```

2. Faça o build e inicie os containers:
   ```
   docker-compose up --build
   ```

3. A API estará disponível em: `http://localhost:8080`

## 📚 Documentação Adicional

Para mais detalhes sobre os endpoints, códigos de resposta e exemplos de uso, consulte nossa [documentação completa](link_para_documentacao).

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor, leia o [guia de contribuição](link_para_guia_de_contribuicao) para saber como começar.

## 📄 Licença

Este projeto está licenciado sob a [MIT License](link_para_licenca).
