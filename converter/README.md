## converter

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
### Working with `.proto` file

Both microservices should use the same [`index.proto`](./grpc/index.proto) file

If any change was made to [`index.proto`](./grpc/index.proto) file, [`index_grpc.pb.go`](./grpc/index_grpc.pb.go) and [`index.pb.go`](./grpc/index.pb.go) files should be regenerated to reflect all of the changes made to `index.proto` (these files shuold not be changed manually)

These files can be regenerated with the [`proto.sh`](./proto.sh) script:

```shell script
bash proto.sh
```

// TODO: dependencies

### Launch

```shell script
# Without AIR
export CGO_CFLAGS_ALLOW='-Xpreprocessor' && go run ./

# With AIR
export CGO_CFLAGS_ALLOW='-Xpreprocessor' && air
```
