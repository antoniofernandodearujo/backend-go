# Sensedia Challenge API

## Languages

- [English](#getting-started)
- [Português Brasileiro](#começando)
- [Español](#comenzando)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

The aplication swagger documentation is available at http://{YOUR_HOST}:${YOUR_PORT}/swagger/index.html

### Prerequisites

You need to have installed the following software:

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

### Installing

After installing the prerequisites, you need to install the dependencies of the project.

```bash
pip install load_dotenv
pip install psycopg2
```

After installing the dependencies, you need to build the database container. Make sure you have your .env file with the correct values, as exemplified in the .env.sample.

```bash
 docker compose up -d
```

After building the database container, you need to run the migrations.

```bash
make run_migrations
```

After running the migrations, populate the database with fake data

```bash
make populate_db
```

After that, you need to build and run the application.

```bash
make build
make run
```

You can also run the application with live reload using air or the make command provided below in the Makefile section.

## MakeFile

```bash
make build
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application. If you already have air installed but our command isn't working, you can run it directly with air.

```bash
make watch
```

Run the test suite

```bash
make test
```

Clean up binary from the last build

```bash
make clean
```

Run the migrations

```bash
make run_migrations
```

Run the migrations rollback, you need to run this once for every migration you want to rollback.

```bash
make rollback_migrations
```

Populate the database with fake data

```bash
make populate_db
```

## Começando

Estas instruções vão te ajudar a obter uma cópia do projeto e executá-lo na sua máquina local para fins de desenvolvimento e teste. Veja as notas sobre implantação para saber como implantar o projeto em um sistema ao vivo.

A documentação do Swagger da aplicação está disponível em http://{SEU_HOST}:{SUA_PORTA}/swagger/index.html

### Pré-requisitos

Você precisa ter instalado o seguinte software:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [Python3](https://www.python.org/downloads/)
- [Pip](https://pip.pypa.io/en/stable/installation/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Rustup](https://rustup.rs/)
- [Sqlx](https://github.com/launchbadge/sqlx/blob/main/sqlx-cli/README.md)

### Instalação

Após instalar os pré-requisitos, você precisa instalar as dependências do projeto.

```bash
pip install load_dotenv
pip install psycopg2
```

Após instalar as dependências, você precisa construir o contêiner do banco de dados. Certifique-se de ter seu arquivo .env com os valores corretos, conforme exemplificado no .env.sample.

```bash
 docker compose up -d
```

Após construir o contêiner do banco de dados, você precisa executar as migrações.

```bash
make run_migrations
```

Após executar as migrações, popule o banco de dados com dados falsos

```bash
make populate_db
```

Depois disso, você precisa buildar e executar a aplicação.

```bash
make build
make run
```

Você também pode executar a aplicação com recarga ao vivo usando air ou o comando make fornecido abaixo na seção Makefile.

## Makefile

```bash
make build
```

Construir a aplicação

```bash
make build
```

Executar a aplicação

```bash
make run
```

Recarga ao vivo da aplicação. Se você já tem air instalado, mas nosso comando não estiver funcionando, você pode executá-lo diretamente com air.

```bash
make watch
```

Executar a suíte de testes

```bash
make test
```

Limpar binário da última construção

```bash
make clean
```

Executar as migrações

```bash
make run_migrations
```

Executar o rollback das migrações, você precisa executar isso uma vez para cada migração que deseja reverter.

```bash
make rollback_migrations
```

Preencher o banco de dados com dados falsos

```bash
make populate_db
```

## Comenzando

Estas instrucciones le ayudarán a obtener una copia del proyecto y a ejecutarlo en su máquina local para propósitos de desarrollo y prueba. Consulte las notas sobre la implementación para saber cómo desplegar el proyecto en un sistema en vivo.

La documentación de Swagger de la aplicación está disponible en http://{TU_HOST}:{TU_PUERTO}/swagger/index.html

### Prerrequisitos

Necesita haber instalado el siguiente software:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [Python3](https://www.python.org/downloads/)
- [Pip](https://pip.pypa.io/en/stable/installation/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Rustup](https://rustup.rs/)
- [Sqlx](https://github.com/launchbadge/sqlx/blob/main/sqlx-cli/README.md)

### Instalación

Después de instalar los prerrequisitos, necesita instalar las dependencias del proyecto.

```bash
pip install load_dotenv
pip install psycopg2
```

Después de instalar las dependencias, necesita construir el contenedor de la base de datos. Asegúrese de tener su archivo .env con los valores correctos, como se ejemplifica en el .env.sample.

```bash
 docker compose up -d
```

Después de construir el contenedor de la base de datos, necesita ejecutar las migraciones.

```bash
make run_migrations
```

Después de ejecutar las migraciones, poblar la base de datos con datos falsos

```bash
make populate_db
```

Después de eso, necesita ejecutar la aplicación.

```bash
make build
make run
```

También puede ejecutar la aplicación con recarga en vivo usando air o el comando make proporcionado a continuación en la sección Makefile.

## Makefile

```bash
make build
```

Construir la aplicación

```bash
make build
```

Ejecutar la aplicación

```bash
make run
```

Recarga en vivo de la aplicación. Si ya tiene air instalado pero nuestro comando no funciona, puede ejecutarlo directamente con air.

```bash
make watch
```

Ejecutar el conjunto de pruebas

```bash
make test
```

Limpiar el binario de la última construcción

```bash
make clean
```

Ejecutar las migraciones

```bash
make run_migrations
```

Ejecutar el rollback de las migraciones, necesita ejecutar esto una vez por cada migración que quiera revertir.

```bash
make rollback_migrations
```

Poblar la base de datos con datos falsos

```bash
make populate_db
```

---
