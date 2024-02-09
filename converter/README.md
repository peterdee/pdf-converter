## converter

This microservice converts PDF files into JPEG files and is connected to [`api-interface`](../api-interface) microservice via gRPC

### Environment variables

This microservice requires its own set of environment variables, `.env` file is required, see [`.env.example`](./.env.example) for details

### Deploy (MacOS)

This microservice uses [ImageMagick](https://imagemagick.org/script/download.php) and [Ghostscript](https://ghostscript.com/docs/9.54.0/Install.htm). Both should be installed into the system:

```shell script
brew install ghostscript

brew install imagemagick
```

ImageMagick also requires `pkg-config` to be installed:

```shell script
brew install pkg-config
```
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

Download all of the necessary packages:

```shell script
go mod download
```

Run the application:

```shell script
export CGO_CFLAGS_ALLOW='-Xpreprocessor' && go run ./
```

Also can be launched with [AIR](https://github.com/cosmtrek/air)

By default gRPC server will be launched on http://localhost:50051
