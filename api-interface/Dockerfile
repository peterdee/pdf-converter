FROM golang:alpine3.19 AS builder
WORKDIR /build
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /api

FROM alpine as release
WORKDIR /release
COPY --from=builder /api .
COPY --from=builder /build/assets/favicon.ico ./assets/favicon.ico
ENV ENV_SOURCE env
EXPOSE 8000
ENTRYPOINT ["/release/api"]
