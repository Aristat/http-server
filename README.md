### Generate api.gen.go

```
oapi-codegen -package api -generate types,chi-server,spec -o internal/app/api/api.gen.go openapi.yaml
```

### Run

```
go mod download
go run main.go s
curl localhost:3000/api/v1/products/1
```
