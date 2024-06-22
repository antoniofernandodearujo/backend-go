# Desafio Backend

## Começando

A documentação do Swagger da aplicação está disponível em http://{SEU_HOST}:{SUA_PORTA}/swagger/index.html

### Pré-requisitos

Você precisará dos seguintes softwares instalados:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [Python3](https://www.python.org/downloads/)
- [Pip](https://pip.pypa.io/en/stable/installation/)
- [libpq-dev](https://pypi.org/project/libpq-dev/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Rustup](https://rustup.rs/)
- [Sqlx](https://github.com/launchbadge/sqlx/blob/main/sqlx-cli/README.md)

### Instalação

Primeiro, instale as dependências do projeto:

```bash
pip install load_dotenv
pip install psycopg2
```

### Configuração do Ambiente

Crie um arquivo `.env` a partir do `.env.sample` e configure-o com os valores apropriados:

### Executando a Aplicação com Docker

Após configurar o ambiente, você pode construir e executar os contêineres Docker, garantindo que seu banco de dados e aplicação rodem em ambientes isolados:

```bash
docker-compose up -d --build
```

### Migrações e População do Banco de Dados

Configure o esquema do banco de dados executando as migrações e depois popule o banco com dados iniciais:

```bash
make run_migrations
make populate_db
```

### Construção e Execução da Aplicação

Por fim, construa e execute a aplicação:

```bash
make build
make run
```

## Testando a api
E por fim testes as rotas!