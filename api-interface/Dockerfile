FROM golang:alpine3.19 AS build
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -o /api
EXPOSE 8000
CMD ["/api"]
