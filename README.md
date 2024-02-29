## pdf-splitting

### Deploy

[`splitter`](./splitter/) microservice uses [ImageMagick](https://imagemagick.org/script/download.php) and [Ghostscript](https://ghostscript.com/docs/9.54.0/Install.htm). Both should be installed into the system (if not using Docker):

```shell script
brew install ghostscript

brew install imagemagick
```

ImageMagick also requires `pkg-config` to be installed:

```shell script
brew install pkg-config
```

### Launch

Compile `proto`:

```shell script
bash proto.sh
```

#### `splitter` microservice

```shell script
# Without AIR
export CGO_CFLAGS_ALLOW='-Xpreprocessor' && go run ./

# With AIR
export CGO_CFLAGS_ALLOW='-Xpreprocessor' && air
```
