# Http Server

It's small example http server based on openapi. Not suitable for microservices, it's better to look to the [kit](https://github.com/go-kit/kit) 

# Generate api.gen.go

```
oapi-codegen -package api -generate types,chi-server,spec -o internal/app/api/api.gen.go openapi.yaml
```

# Run

```
go mod download
go run main.go s
curl localhost:3000/api/v1/products/1
```
