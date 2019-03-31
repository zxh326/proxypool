
#build stage
FROM golang:alpine AS builder
ENV GOPROXY https://go.likeli.top
ENV GO111MODULE on
WORKDIR /go/cache
RUN apk --no-cache add git g++
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/release/app /app
ENTRYPOINT ./app
EXPOSE 8008
