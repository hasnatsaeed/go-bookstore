#  base image for Go
FROM golang:latest
# Set the Current Working Directory inside the container

WORKDIR /app

RUN git clone https://github.com/hasnatsaeed/go-bookstore.git

WORKDIR /app/go-bookstore

#RUN go build -o go-bookstore

EXPOSE 8080

ENTRYPOINT ["go", "run", "/app/go-bookstore/cmd/main/main.go"]