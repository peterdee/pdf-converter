FROM golang:alpine3.19 AS build
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN apk --no-cache add msttcorefonts-installer fontconfig && \
 update-ms-fonts && \
 fc-cache -f
RUN apk --update add pkgconfig
RUN apk --update add ghostscript
RUN apk --update add imagemagick
ARG CGO_CFLAGS_ALLOW="-Xpreprocessor"
ENV CGO_CFLAGS_ALLOW="-Xpreprocessor"
RUN go mod download
COPY . .
RUN go get gopkg.in/gographics/imagick.v3
RUN go build -o /converter
EXPOSE 50051
CMD ["/converter"]
