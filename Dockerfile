FROM golang:1.20.2

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD ["/main"]