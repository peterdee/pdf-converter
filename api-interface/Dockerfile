FROM golang:alpine3.19 AS build
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -o api

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
WORKDIR /api
COPY --from=build /api /.env ./
EXPOSE 8000
CMD ["api"]
