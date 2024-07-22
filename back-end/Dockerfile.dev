FROM golang:1.22

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["CompileDaemon", "--build=go build -o main .", "--command=./main"]