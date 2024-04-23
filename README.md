# Hasher

The application for creating "hashes" from client data.
App is divided into three parts and each part works independently of the others.

- REST API
- GRPC API
- RabbitMQ consumer

## Requirements

- docker and docker-compose (required)
- [migrate](https://github.com/golang-migrate/migrate?tab=readme-ov-file) (required)
- makefile (required)
- protoc (required)

## Installation

1. Clone the repository
2. Go to the project directory and change the `.env` file
3. Run the command `make docker_start` to start the application
4. Run the command `make migration_up` to run the migrations

## Usage

In `.env` file you can change the port for each service. And you can change version of the services.
By default, the application runs on ports:

- REST API: 8080
- GRPC API: 8089
- RabbitMQ: 5672

Default version of the services is `http`.
Remember, if id in request is lower than 1, application will not save hash into database.

### REST API

You can use three endpoints:

- `POST /v1/email` - to hash email

Example request:

```json
{
  "id": 0,
  "address": "john@doe.com"
}
```

- `POST /v1/phone` - to hash phone number

Example request:

```json
{
  "id": 0,
  "number": "123456789",
  "country_code": "PL"
}
```

- `POST /v1/name` - to hash name and surname

Example request:

```json
{
  "id": 0,
  "name": "John",
  "surname": "Doe"
}
```

### GRPC API

You can use three methods:

- `HashEmail` - to hash email

Example request:

```json
{
  "id": 0,
  "address": "john@doe.com"
}
```

- `HashPhone` - to hash phone number

Example request:

```json
{
  "id": 0,
  "number": "123456789",
  "country_code": "PL"
}
```

- `HashName` - to hash name and surname

Example request:

```json
{
  "id": 0,
  "name": "John",
  "surname": "Doe"
}
```

### RabbitMQ consumer

Listens to the three queues:

- `email_change` - to hash email

Example request:

```json
{
  "id": 0,
  "address": "john@doe.com"
}
  ```

- `phone_change` - to hash phone number

Example request:

```json
{
  "id": 0,
  "number": "123456789",
  "country_code": "PL"
}
```

- `name_change` - to hash name and surname

Example request:

```json
{
  "id": 0,
  "name": "John",
  "surname": "Doe"
}
```

## Tests

To run tests, use the command `make test`.
Test coverage the most important part of the application (hashing).
All tests are in files with the `_test.go` suffix.

## Deployment

- If u want to deploy the application, you can use the `docker-compose.yml` file. Remember to change the `.env` file.
- If you want to add new migrations, you can use the `make migration_create` command.
- If you want to change something in the proto file, you can use the `make grpc` command to rebuild.
- If you want change or add something into RabbitMQ consumer, go to the `broker,go` file.
- If you want to change or add something into REST API, go to the `server.go` file.
- If you want to change or add something into GRPC API, go to the `server.go` file or `grpc.go` file.
- If you want to add other SQL provider, go to the `database.go` file and `store.go` file.
- If you want add kafka consumer, go to the `broker.go` file.

## Authors

- [Robert Gontarski](https://linkedin.com/in/robert-gontarski)

## License

This project is licensed under the [MIT License](LICENSE).