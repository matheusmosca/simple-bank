<h1 align="center">:moneybag: Simple Bank :moneybag:</h1>

<p align="center">Simple Bank is a simple REST API that allows users to perform transferences with each other.</p>

## :wrench: Technologies

- Golang
- Docker
- PostgreSQL

## :books: Libraries

- [Gorilla Mux](https://github.com/gorilla/mux)
- [PQ](https://github.com/lib/pq)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)
- [Validator](https://github.com/go-playground/validator)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Google uuid](https://github.com/google/uuid)
- [Golang-migrate](https://github.com/golang-migrate/migrate)
- [Go-bindata](https://github.com/kevinburke/go-bindata)
- [Testify](https://github.com/stretchr/testify)
- [Moq](https://github.com/matryer/moq)

# :triangular_flag_on_post: Endpoints

## `/accounts`

### `/api/v1/accounts - POST` Creates a new account. Example of request body:

```json
{
  "name": "Maria",
  "secret": "12345678",
  "cpf": "103.913.350-90"
}
```

### `/api/v1/accounts - GET` Fetch all accounts

### `/api/v1/accounts/{account_id}/balance - GET` Get the account's balance

### Details

- Accounts endpoints does not require authentication
- New accounts starts with 0 balance
- The account's balance is always integer
- The `cpf` field should match `xxx.xxx.xxx-xx`. Take a look at [this website](https://www.4devs.com.br/gerador_de_cpf) for generate random cpfs.
- The `secret` field must have between 6 and 50 characters

## `/login`

### `/api/v1/login - POST` Creates a JWT token. Example:
```json
// Request body
{
	"cpf": "084.792.650-86",
	"secret": "12345678"
}
// Response
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAsIklkIjoiZGY2YWNlODktNGE0Yy00NGY5LTk3OGMtNTIxMTEzNDEwMDM1IiwiSXNzdWVkQXQiOjE2MjA1MTE0Mjl9.v1kzBo3GdIXO1vhTq9_icIgfdgE7981KZ5CaQlw3Bvs"
}
```

## `/transfers`

### `/api/v1/transfers - POST` Perform a transference between two accounts. Example:

```json
// Request body
{
  "account_destination_id": "eb7b34eb-643b-4e2b-9e81-7641b3e45327",
  "amount": 100
}
// Response body
{
  "id": "6e5d9213-2c86-4cb8-bfc7-ddafd9237d34",
  "account_destination_id": "eb7b34eb-643b-4e2b-9e81-7641b3e45327",
  "account_origin_id": "df6ace89-4a4c-44f9-978c-521113410035",
  "amount": 100,
  "created_at": "2021-05-08T22:06:18.439343Z"
}
```

### `/api/v1/transfers - GET` Get the transferences of the authenticated user

### Details

- Bearer token is required to perform requests
- The `account_origin_id` comes from the authenticated user, and can't be equal the `account_destination_id` (make a transference to yourself does not make sense)
- The authenticated user must have sufficient funds to perform a transference

# :checkered_flag: How to run (needs Docker :whale:)

- Copy the `.env.example` file content into a new file called `.env`
- After using `make dev-local` or `make dev-docker`, the API will be available in `http://localhost:3000/api/v1`

### Running the application locally using Go and Docker

```zsh
  make dev-local
```

### Running with Docker (Go not required)
```zsh
  make dev-docker
```

## :vertical_traffic_light: Testing

### Run tests
```zsh
  make test
```

### See test coverage
```zsh
  make test-coverage
```
