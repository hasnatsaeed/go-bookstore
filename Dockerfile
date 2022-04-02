#  base image for Go
FROM golang:latest

RUN mkdir /app

# Set the Current Working Directory inside the container
WORKDIR /app

ADD . /app

# Build the Go app
RUN go get github.com/hasnatsaeed/go-bookstore && go build -o app/cmd/main.go

ENTRYPOINT ["/app/cmd/main"]