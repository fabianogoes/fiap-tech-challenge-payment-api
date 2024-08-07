# FIAP challenge - Payment

[![CI](https://github.com/fabianogoes/fiap-tech-challenge-payment-api/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/fabianogoes/fiap-tech-challenge-payment-api/actions/workflows/ci-cd.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=fabianogoes_fiap-tech-challenge-payment-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=fabianogoes_fiap-tech-challenge-payment-api)
[![Sonar Coverage](https://sonarcloud.io/api/project_badges/measure?project=fabianogoes_fiap-tech-challenge-payment-api&metric=coverage)](https://sonarcloud.io/summary/new_code?id=fabianogoes_fiap-tech-challenge-payment-api)
<img src="https://sonarcloud.io/images/project_badges/sonarcloud-white.svg" alt="Scanned on SonarCloud" height="20px" />

## Project Architecture by Clean Architecture

- `app/web`: diretório para os principais pontos de entrada, injeção dependência ou comandos do aplicativo. O subdiretório ‘web’ contém o ponto de entrada principal a API REST.
- `domain/entities`: diretório que contém modelos/entidades de domínio que representam os principais conceitos de negócios.
- `domain/usecases`: diretório que contém Serviços de Domínio ou Use Cases.
- `domain/ports`: diretório que contém ‘interfaces’ ou contratos definidos que os adaptadores devem seguir.
- `frameworks/rest`: diretório que contém os controllers e manipulador de requisições REST.
- `frameworks/rest/dto`: diretório que contém objetos/modelo de request e response.
- `frameworks/repository`: diretório que contém adaptadores de banco de dados exemplo para PostgreSQL.
- `frameworks/repository/dbo`: diretório que contém objetos/entidades de banco de dados.
- `.infra`: diretório que contém arquivos de infrainstrutura
- `.infra/kubernetes`: diretório que contém os manifestos kubernetes

## Stack

- [x] [Go][0]
- [x] [Domain-Driven Design][6]
- [x] [Hexagonal Architecture][5]
- [x] [Gin Web Framework][1] - Routes, JSON validation, Error management, Middleware support
- [x] [MongoDB][3] - Database persistence
- [x] [GORM ORM library for Golang][2]
- [x] [Slog](https://pkg.go.dev/log/slog) - Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs. 
- [x] [GoDotEnv](https://github.com/joho/godotenv) - A Go (golang) port of dotenv project (which loads env vars from a .env file).

## Issues

- [ ] [gin-swagger](https://github.com/swaggo/gin-swagger) - gin middleware to automatically generate RESTful API documentation with Swagger 2.0.
- [ ] [swag](https://github.com/swaggo/swag) - Swag converts Go annotations to Swagger Documentation 2.0
- [ ] [CORS gin's middleware](https://github.com/gin-contrib/cors) - Gin middleware/handler to enable CORS support.

## Development

Dependencies

- [Go Installation](https://go.dev/doc/install)

Check for go version 1.21.3

```shell
go version
```

Preparing app

```shell
git clone git@github.com:fabianogoes/fiap-tech-challenge-payment-api.git
cd fiap-tech-challenge-payment-api
go mod tidy
````

### Running

```shell
docker-compose up -d mongo && go run app/web/main.go
```

## Testing using Docker/Docker Compose

```shell
docker-compose up -d

curl --request GET --url http://localhost:8010/health

## response 
{"status":"UP"}
```

## Docker Commands

```shell
docker login -u=fabianogoes
docker build -t fabianogoes/payment-api:latest .
docker tag fabianogoes/payment-api:latest fabianogoes/payment-api:latest
docker push fabianogoes/payment-api:latest
```

## Run test

```shell
go test -v ./...
```

## Run test with coverage

```shell
 go test -coverprofile=coverage.out ./... &&  go tool cover -func=coverage.out
```

[0]: https://go.dev/
[1]: https://gin-gonic.com/
[2]: https://gorm.io/index.html
[3]: https://www.mongodb.com/
[5]: https://alistair.cockburn.us/hexagonal-architecture/
[6]: https://www.amazon.com/dp/0321125215?ref_=cm_sw_r_cp_ud_dp_0M66DHP14SJ5GBBJCRNP
