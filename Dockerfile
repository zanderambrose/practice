FROM golang:1.17

WORKDIR /app
COPY . /app

RUN go build -o main /app/api/main.go

EXPOSE 8080

CMD ["./main"]
