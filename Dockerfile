FROM golang:1.23.1

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

RUN go build -o main main.go

CMD ["./main"]