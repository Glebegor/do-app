FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x postgres-db.sh

RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]