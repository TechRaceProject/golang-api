FROM golang:1.21.6 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM golang:1.21.6

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/src ./src

EXPOSE 8000

CMD ["./main"]
