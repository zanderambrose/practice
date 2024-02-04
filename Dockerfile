FROM golang:1.17

WORKDIR /app
COPY . /app

RUN go build -o main /app/cmd/api/main.go

EXPOSE 8080

CMD ["./main"]
