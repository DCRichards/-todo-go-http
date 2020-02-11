FROM golang:1.13-alpine

RUN apk update && apk add git

WORKDIR /go/src/github.com/dcrichards/todo-go-http

RUN go get -u github.com/cosmtrek/air

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["air", "-c", ".air.conf"]
