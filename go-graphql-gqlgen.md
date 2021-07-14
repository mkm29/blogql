# Building a GraphQL API in Go using gqlgen
A simple GraphQL API that implements [this specification](./src/cmd/go-graphql/schema.graphql) for a To Do list
Refer to the medium article [Building a GraphQL API in Go using gqlgen](https://medium.com/weareservian/building-a-graphql-api-in-go-using-gqlgen-f7a42eba2193)
for more details about this project.

## Running locally
```
docker-compose up
cd src/cmd/go-graphql
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init --verbose
POSTGRES_USER=graphql_api POSTGRES_PASSWORD=_please_change_me_now_ POSTGRES_URL=localhost POSTGRES_DB=graphql go run server.go
```