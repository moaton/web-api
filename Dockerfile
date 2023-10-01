FROM golang:latest
RUN apt-get update
WORKDIR /go/src/app
COPY . .
RUN go build -o web-api ./cmd/main.go
CMD ['./web-api']