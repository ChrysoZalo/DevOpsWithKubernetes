FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go layout.html ./

RUN CGO_ENABLE=0 GOOS=linux go build -o /todo-app

CMD ["/todo-app"]