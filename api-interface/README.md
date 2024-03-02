## api-interface

This microservice provides necessary REST APIs via HTTP and it's connected to [`converter`](../converter) microservice via gRPC

### Environment variables

This microservice requires its own set of environment variables, `.env` file is required, see [`.env.example`](./.env.example) for details

### Working with `.proto` file

Both microservices should use the same [`index.proto`](./grpc/index.proto) file

If any change was made to [`index.proto`](./grpc/index.proto) file, both [`index_grpc.pb.go`](./grpc/index_grpc.pb.go) and [`index.pb.go`](./grpc/index.pb.go) files should be regenerated to reflect all of the changes made to `index.proto` (these files shuold not be changed manually)

These files can be regenerated with the [`proto.sh`](./proto.sh) script:

```shell script
bash proto.sh
```

// TODO: dependencies

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

### Available APIs

// TODO: list APIs or create a separate file
