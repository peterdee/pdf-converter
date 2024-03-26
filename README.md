## pdf-converter

Convert PDF files into JPEG images using Golang

This repository contains 2 microservices: [`api-interface`](./api-interface) and [`converter`](./converter)

`api-interface` provides REST APIs via HTTP to queue PDF conversion and get the resulting archive

`converter` converts queued PDF files and returns ZIP-archive after PDF has been converted

Microservices are connected via gRPC, conversion queue is implemented using MongoDB, PDF files are converted into JPEG using ImageMagick

### Deploy

Clone the repository:

```shell script
git clone https://github.com/peterdee/pdf-converter
cd ./pdf-converter
gvm use 1.22.0
```

Additional deployment and launching instructions are available in [api-interface/README.md](./api-interface/README.md) and [converter/README.md](./converter/README.md)

### Launch everything with Docker Compose

Everything can be launched with Docker Compose ([docker-compose.yml](./docker-compose.yml))

Create `.env` files for `api-interface` and `converter` microservices before the launch


```shell script
docker compose up
```

`api-interface` application will be available at http://localhost:8000

### License

[MIT](./LICENSE.md)
