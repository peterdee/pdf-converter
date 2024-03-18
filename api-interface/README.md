## api-interface

This microservice provides necessary REST APIs via HTTP and it's connected to [`converter`](../converter) microservice via gRPC

### Environment variables

This microservice requires its own set of environment variables, `.env` file is required, see [`.env.example`](./.env.example) for details

### Working with `.proto` file (MacOS)

Both microservices should use the same [`index.proto`](./grpc/index.proto) file

If any change was made to [`index.proto`](./grpc/index.proto) file, [`index_grpc.pb.go`](./grpc/index_grpc.pb.go) and [`index.pb.go`](./grpc/index.pb.go) files should be regenerated to reflect all of the changes made to `index.proto` (these files shuold not be changed manually)

Install all of the necessary dependencies for `.proto` file compilation:

```shell script
# Protocol Buffer compiler
brew install protobuf

# Golang wrappers
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Compile [`index.proto`](./grpc/index.proto) file using [`proto.sh`](./proto.sh) script:

```shell script
bash proto.sh
```

### Launch

Download all of the necessary packages:

```shell script
go mod download
```

Run the application:

```shell script
go run ./
```

Also can be launched with [AIR](https://github.com/cosmtrek/air)

HTTP server will be available at http://localhost:8000

Alternatively use Docker:

```shell script
docker run -p 8000:8000 -it $(docker build -q .)
```

### Available APIs

Available APIs are listed in [APIS.md](./APIS.md)
