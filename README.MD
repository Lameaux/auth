# Auth

Authentication service.

http://euromoby.com

### Running

```
go run cmd/main.go
```

### ENV variables

- PORT (default `8080`)
- WAIT_TIMEOUT (default `15`)
- DATABASE_URI (default `postgres://root:heslo@localhost:5432/auth_test?&pool_max_conns=10`)

### Testing

```
go test -v ./...
```

### Migrations

https://github.com/pressly/goose

```
cd migrations
goose postgres "user=root password=heslo dbname=auth_test" reset
goose postgres "user=root password=heslo dbname=auth_test" up
```

### Endpoints

#### Create User

- POST /v1/users

