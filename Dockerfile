FROM golang:latest

WORKDIR /app

COPY src/go.* ./
RUN go mod download

COPY src/ .

RUN go build -o main .

CMD ["./main"]
