## pdf-splitting

### Deploy

[`splitter`](./splitter/) microservice uses [ImageMagick](https://imagemagick.org/script/download.php) and [Ghostscript](https://ghostscript.com/docs/9.54.0/Install.htm). Both should be installed into the system (if not using Docker):

```shell script
brew install ghostscript

brew install imagemagick
```

### Launch

Compile `proto`:

```shell script
bash proto.sh
```
