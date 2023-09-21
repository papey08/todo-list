FROM golang:1.21

ENV GOPATH=/

WORKDIR /go/src/todo-list
COPY . .

RUN go mod download
RUN go build -o todo-list-app cmd/server/main.go

CMD ["./todo-list-app"]
